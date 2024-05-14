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
	memberOfFromLdap := result.Entries[0].GetAttributeValues("memberOf")
	log.Printf("name: %v\n", nameFromLdap)
	log.Printf("memberOf: %v\n", memberOfFromLdap)

	for _, group := range memberOfFromLdap {
		if group == h.ldapGroup {
			log.Printf("Authorized")
		}
	}

	// db := c.Get(h.dbKey).(*gorm.DB)
	// var user struct {
	// 	Id       int
	// 	Username string
	// 	Password string
	// 	Enabled  int
	// }
	// _ = db.Raw("SELECT id, username, password, enabled FROM users WHERE email = ?;", email).Scan(&user)

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = 123
	session.Values["user_email"] = email
	session.Values["username"] = "LDAPteste"

	if remember == "true" {
		session.Options.MaxAge = 84600 * 30
	}

	if err := session.Save(c.Request(), c.Response()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(http.StatusSeeOther)

}
