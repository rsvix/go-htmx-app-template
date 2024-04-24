package registerhandler

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/rsvix/go-htmx-app-template/internal/emails"
	"github.com/rsvix/go-htmx-app-template/internal/hash"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type postRegisterHandlerParams struct {
}

func PostRegisterHandler() *postRegisterHandlerParams {
	return &postRegisterHandlerParams{}
}

func (h *postRegisterHandlerParams) Serve(c echo.Context) error {

	email := c.Request().FormValue("email")
	username := c.Request().FormValue("username")
	firstname := c.Request().FormValue("firstname")
	lastname := c.Request().FormValue("lastname")
	password := c.Request().FormValue("password")
	password_conf := c.Request().FormValue("passwordconf")
	// log.Printf("email: %s\nfirstname: %s\nlastname: %s\npassword: %s\npassword_conf: %s\n", email, firstname, lastname, password, password_conf)

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid email</h2>")
	}

	if !utils.IsValidName(username) {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid username</h2>")
	}

	if !utils.IsValidName(firstname) {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid first name</h2>")
	}

	if !utils.IsValidName(lastname) {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid lastname</h2>")
	}

	if password != password_conf {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Passwords do not match</h2>")
	}

	if !utils.IsValidPasswordV1(password) {
		// if !utils.IsValidPasswordV2(password) {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid password</h2>")
	}

	// hashPassword, err := hash.HashPasswordV1(password)
	hashPassword, err := hash.HashPasswordV2(password)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}
	// log.Printf("hashPassword: %s\n", hashPassword)

	// ip := c.RealIP()
	ip, ipErr := utils.GetIP(c.Request())
	if ipErr != nil {
		log.Printf("Error obtaining ip from request: %v\n", ipErr)
	}

	user := structs.User{
		Email:     email,
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
		Password:  hashPassword,
		// Activationtoken:               activationToken,
		// Activationtokenexpiration:     timeExp,
		// Passwordchangetokenexpiration: timeNow,
		Registerip: ip,
	}

	db := c.Get("__db").(*gorm.DB)
	// var ID int64
	// db.Raw("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id", email, hash_password).Scan(&ID)
	// log.Printf("User %v created with ID: %v\n", email, ID)
	result := db.Create(&user)
	if err := result.Error; err != nil {
		log.Println(err.Error())
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}
	// log.Printf("Create user result: %s\n", result)

	id := strconv.FormatInt(int64(user.ID), 10)

	activationToken, err := hash.GenerateToken(true, id)
	if err != nil {
		log.Printf("Error generating activation token: %v\n", err)
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}

	timeNow := time.Now().UTC()
	timeExp := timeNow.Add(24 * time.Hour)

	var emailResp string
	result2 := db.Raw("UPDATE users SET activationtoken = ?, activationtokenexpiration = ?, passwordchangetokenexpiration = ? WHERE id = ? RETURNING email;", activationToken, timeExp, timeNow, id).Scan(&emailResp)
	log.Printf("emailResp: %v\n", emailResp)
	log.Printf("Query result: %v\n", result2)

	appPort, _ := os.LookupEnv("APP_PORT")
	activationUrl := fmt.Sprintf("http://localhost:%s/tkn/%s", appPort, activationToken)
	// log.Printf("Activation Url: %s\n", activationUrl)

	// Must configure SMTP server or other email sending service
	if _, ok := os.LookupEnv("SENDER_PSWD"); ok {
		if err := emails.SendActivationMail(email, activationUrl, emails.DefaultParams()); err != nil {
			log.Printf("Error sending email: %s\n", err)
		}
		msg := fmt.Sprintf("<h2>Activation email sent to<br/>%s</h2>", email)
		return c.HTML(http.StatusOK, msg)
	}
	// For testing without email sender configured
	msg := fmt.Sprintf("<h2><div style=\"text-decoration-line: underline;\"><a href=\"%s\">Activate your account<br/>by clicking here</a></div></h2>", activationUrl)
	return c.HTML(http.StatusOK, msg)
}
