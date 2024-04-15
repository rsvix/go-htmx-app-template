package handlers

import (
	"encoding/hex"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type GetProcessHandler struct {
}

func NewGetProcessHandler() *GetProcessHandler {
	return &GetProcessHandler{}
}

func (i GetProcessHandler) ServeHTTP(c echo.Context) error {

	token := c.Param("token")
	log.Printf("token: %s\n", token)

	if !strings.Contains(token, "O") {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	// Check if session is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		validator := strings.Split(token, "O")[0]
		extraInfo := strings.Split(token, "O")[1]
		log.Printf("validator: %s\n", validator)
		log.Printf("extraInfo: %s\n", extraInfo)

		decoded, err := hex.DecodeString(extraInfo)
		if err != nil {
			log.Println(err)
			return err
		}
		decodedStr := string(decoded[:])
		mode := strings.Split(decodedStr, "@")[0]
		id := strings.Split(decodedStr, "@")[1]
		log.Printf("mode: %s\n", mode)
		log.Printf("id: %s\n", id)

		if mode == "activate" {
			log.Println("Processing activation")
		} else if mode == "resetpwd" {
			log.Println("Processing password reset")
		}
		return nil
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
