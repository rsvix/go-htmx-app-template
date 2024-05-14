package ldaploginhandler

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
	"gorm.io/gorm"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type postLdapLoginHandlerParams struct {
	dbKey     string
	baseDn    string
	ldapGroup string
}

func PostLdapLoginHandler() *postLdapLoginHandlerParams {
	return &postLdapLoginHandlerParams{
		dbKey:     os.Getenv("DB_CONTEXT_KEY"),
		baseDn:    os.Getenv("LDAP_BASE_DN"),
		ldapGroup: os.Getenv("LDAP_GROUP"),
	}
}

func (h *postLdapLoginHandlerParams) Serve(c echo.Context) error {

	email := c.Request().FormValue("email")
	password := c.Request().FormValue("password")
	remember := c.Request().FormValue("remember")

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid credentials")
	}

	ldapUser := strings.Split(email, "@")[0]

	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/
	ldapURL := os.Getenv("LDAP_URL")
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Printf("LDAP conn error: %v\n", err)
		return c.HTML(http.StatusUnprocessableEntity, "Error, please try again")
	}
	defer l.Close()

	// Upgrade to TLS connection
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Printf("LDAPS conn error: %v\n", err)
		return c.HTML(http.StatusUnprocessableEntity, "Error, please try again")
	}

	err = l.Bind(email, password)
	if err != nil {
		log.Printf("LDAP bind error: %v\n", err)
		return c.HTML(http.StatusUnprocessableEntity, "Invalid credentials")
	}
	log.Println("LDAP bind successfull")

	searchReq := ldap.NewSearchRequest(
		h.baseDn,
		ldap.ScopeWholeSubtree,
		0,
		0,
		0,
		false,
		fmt.Sprintf("(sAMAccountName=%s)", ldap.EscapeFilter(ldapUser)),
		[]string{"*"},
		[]ldap.Control{},
	)

	result, err := l.Search(searchReq)
	if err != nil {
		log.Printf("failed to query LDAP: %v\n", err)
		return c.HTML(http.StatusUnprocessableEntity, "Error, please try again")
	}
	log.Println("Got", len(result.Entries), "search results")
	// result.PrettyPrint(2)

	nameFromLdap := result.Entries[0].GetAttributeValues("name")[0]
	splitNameFromLdap := strings.Split(nameFromLdap, " ")
	memberOfFromLdap := result.Entries[0].GetAttributeValues("memberOf")
	// log.Printf("name: %v\n", nameFromLdap)
	// log.Printf("memberOf: %v\n", memberOfFromLdap)

	for _, group := range memberOfFromLdap {
		if group == h.ldapGroup {
			log.Printf("Authorized - User in group")

			ip, ipErr := utils.GetIP(c.Request())
			if ipErr != nil {
				log.Printf("Error obtaining ip from request: %v\n", ipErr)
				ip = ""
			}

			db := c.Get(h.dbKey).(*gorm.DB)

			// RAW
			var ID int
			result := db.Raw("INSERT INTO users (email, username, firstname, lastname, registerip, enabled) VALUES (?, ?, ?, ?, ?, ?) RETURNING id;",
				email,
				ldapUser,
				splitNameFromLdap[0],
				splitNameFromLdap[len(splitNameFromLdap)-1],
				ip,
				1,
			).Scan(&ID)

			if err := result.Error; err != nil {
				if strings.Contains(err.Error(), "violates unique constraint \"users_email_key\"") {
					res2 := db.Raw("UPDATE users SET lastip = ? WHERE email = ? RETURNING id;", ip, email).Scan(&ID)
					if err := res2.Error; err != nil {
						log.Printf("err: %v\n", err)
						return c.HTML(http.StatusInternalServerError, "An error occured<br>please try again")
					}
				} else {
					log.Printf("err: %v\n", err)
					return c.HTML(http.StatusInternalServerError, "An error occured<br>please try again")
				}
			}
			// log.Printf("id: %v\n", ID)

			session, err := session.Get("authenticate-sessions", c)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			session.Values["authenticated"] = true
			session.Values["user_id"] = ID
			session.Values["user_email"] = email
			session.Values["username"] = ldapUser
			session.Values["ldap"] = true

			if remember == "true" {
				session.Options.MaxAge = 84600 * 30
			}

			if err := session.Save(c.Request(), c.Response()); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			c.Response().Header().Set("HX-Redirect", "/")
			return c.NoContent(http.StatusSeeOther)
		}
	}
	log.Printf("Not authorized - User not in group")
	return c.HTML(http.StatusUnprocessableEntity, "User not in access group")
}
