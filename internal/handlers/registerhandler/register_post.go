package registerhandler

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rsvix/go-htmx-app-template/internal/emails"
	"github.com/rsvix/go-htmx-app-template/internal/hash"
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

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid email")
	}

	if !utils.IsValidUsername(username) {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid username")
	}

	if !utils.IsValidName(firstname) {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid firstname")
	}

	if !utils.IsValidName(lastname) {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid lastname")
	}

	if password != password_conf {
		return c.HTML(http.StatusUnprocessableEntity, "Passwords do not match")
	}

	if !utils.IsValidPasswordV1(password) {
		// if !utils.IsValidPasswordV2(password) {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid password")
	}

	// hashPassword, err := hash.HashPasswordV1(password)
	hashPassword, err := hash.HashPasswordV2(password)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return c.HTML(http.StatusInternalServerError, "An error occured<br>please try again")
	}

	ip, ipErr := utils.GetIP(c.Request())
	if ipErr != nil {
		log.Printf("Error obtaining ip from request: %v\n", ipErr)
		return c.HTML(http.StatusInternalServerError, "An error occured<br>please try again")
	}

	db := c.Get("__db").(*gorm.DB)

	// RAW
	var ID int
	result := db.Raw("INSERT INTO users (email, username, firstname, lastname, password, registerip) VALUES (?, ?, ?, ?, ?, ?) RETURNING id;",
		email,
		username,
		firstname,
		lastname,
		hashPassword,
		ip,
	).Scan(&ID)
	id := strconv.FormatInt(int64(ID), 10)

	// // GORM
	// user := structs.User{
	// 	Email:      email,
	// 	Username:   username,
	// 	Firstname:  firstname,
	// 	Lastname:   lastname,
	// 	Password:   hashPassword,
	// 	Registerip: ip,
	// }
	// result := db.Create(&user)
	// id := strconv.FormatInt(int64(user.ID), 10)

	if err := result.Error; err != nil {
		log.Printf("Error creating user in database: %s\n", err.Error())
		if strings.Contains(err.Error(), "violates unique constraint \"users_email_key") {
			return c.HTML(http.StatusInternalServerError, "Email already registered")
		} else if strings.Contains(err.Error(), "violates unique constraint \"users_username_key") {
			return c.HTML(http.StatusInternalServerError, "Username not available")
		}
		return c.HTML(http.StatusInternalServerError, "An error occured<br>please try again")
	}

	activationToken, err := hash.GenerateToken(true, id)
	if err != nil {
		log.Printf("Error generating activation token: %v\n", err)
		return c.HTML(http.StatusInternalServerError, "An error occured<br>please try again")
	}

	timeNow := time.Now().UTC()
	timeExp := timeNow.Add(24 * time.Hour)

	// RAW
	db.Exec("UPDATE users SET activationtoken = ?, activationtokenexpiration = ?, passwordchangetokenexpiration = ? WHERE id = ?;",
		activationToken,
		timeExp,
		timeNow,
		id,
	)
	// GORM
	// db.Table("users").Where("WHERE id = ?", id).Updates(map[string]interface{}{"activationtoken": activationToken, "activationtokenexpiration": timeExp, "passwordchangetokenexpiration": timeNow})

	appPort, _ := os.LookupEnv("APP_PORT")
	activationUrl := fmt.Sprintf("http://localhost:%s/activation?%s", appPort, activationToken)

	// Must configure SMTP server or other email sending service
	if _, ok := os.LookupEnv("SENDER_PSWD"); ok {
		if err := emails.SendActivationMail(email, activationUrl, emails.DefaultParams()); err != nil {
			log.Printf("Error sending email: %s\n", err)
		}
		msg := fmt.Sprintf("Activation email sent to<br/>%s", email)
		return c.HTML(http.StatusOK, msg)
	}
	// For testing without email sender configured
	msg := fmt.Sprintf("<div style=\"text-decoration-line: underline;\"><a href=\"%s\">Activate your account<br/>by clicking here</a></div>", activationUrl)
	return c.HTML(http.StatusOK, msg)
}
