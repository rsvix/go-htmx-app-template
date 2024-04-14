package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/rsvix/go-htmx-app-template/internal/hash"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostRegisterHandler struct {
}

type PostRegisterHandlerParams struct {
}

func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{}
}

func (h *PostRegisterHandler) ServeHTTP(c echo.Context) error {

	email := c.Request().FormValue("email")
	firstname := c.Request().FormValue("firstname")
	lastname := c.Request().FormValue("lastname")
	password := c.Request().FormValue("password")
	password_conf := c.Request().FormValue("passwordconf")
	// log.Printf("email: %s\nfirstname: %s\nlastname: %s\npassword: %s\npassword_conf: %s\n", email, firstname, lastname, password, password_conf)

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid email</h2>")
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

	// hashPassword, err := hash.HashPassword(password)
	hashPassword, err := hash.NewHPasswordHash().HashPasswordV2(password)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}
	// log.Printf("hashPassword: %s\n", hashPassword)

	activationToken, err := hash.GenerateActivationToken()
	if err != nil {
		log.Printf("Error generating activation token: %v\n", err)
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}
	// log.Printf("activationToken: %v\n", activationToken)

	// ip := c.RealIP()
	ip, ipErr := utils.GetIP(c.Request())
	if ipErr != nil {
		log.Printf("Error obtaining ip from request: %v\n", ipErr)
	}

	timeNow := time.Now().UTC()
	timeExp := timeNow.Add(24 * time.Hour)

	user := structs.User{
		Email:                         email,
		Firstname:                     firstname,
		Lastname:                      lastname,
		Password:                      hashPassword,
		Activationtoken:               activationToken,
		Activationtokenexpiration:     timeExp,
		Passwordchangetokenexpiration: timeNow,
		Registerip:                    ip,
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

	maskedId := strconv.FormatUint(uint64((user.ID + 575426791)), 10)
	appPort, _ := os.LookupEnv("APP_PORT")
	activationUrl := fmt.Sprintf("http://localhost:%s/activation/%s/%s", appPort, maskedId, activationToken)
	// log.Printf("Activation Url: %s\n", activationUrl)

	// Must configure SMTP server or other email sending service
	// mail.SendActivationEmail(email, activationUrl)
	// msg := fmt.Sprintf("<h2>Activation email sent to<br/>%s</h2>", email)
	// return c.HTML(http.StatusOK, msg)

	// For testing without email sender configured
	msg := fmt.Sprintf("<h2><div style=\"text-decoration-line: underline;\"><a href=\"%s\">Activate your account<br/>by clicking here</a></div></h2>", activationUrl)
	return c.HTML(http.StatusOK, msg)
}
