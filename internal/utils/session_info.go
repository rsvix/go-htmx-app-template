package utils

import (
	"errors"
	"log"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type SessionInfo struct {
	Id       int
	Email    string
	Username string
}

func GetSessionInfo(c echo.Context) (SessionInfo, error) {

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return SessionInfo{}, errors.New("couldn't get session info")
	}

	var id int
	if value, ok := session.Values["user_id"].(int); ok {
		id = value
	} else {
		return SessionInfo{}, errors.New("couldn't get session info")
	}

	var email string
	if value, ok := session.Values["user_email"].(string); ok {
		email = value
	} else {

		return SessionInfo{}, errors.New("couldn't get session info")
	}

	var username string
	if value, ok := session.Values["username"].(string); ok {
		username = value
	} else {
		return SessionInfo{}, errors.New("couldn't get session info")
	}

	s := SessionInfo{
		Id:       id,
		Email:    email,
		Username: username,
	}

	return s, nil
}
