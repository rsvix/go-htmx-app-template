package utils

import (
	"errors"
	"log"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type SessionInfo struct {
	Authenticated bool
	Id            int
	Email         string
	Username      string
}

func GetSessionInfo(c echo.Context) (SessionInfo, error) {

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Println("error getting session")
		return SessionInfo{}, errors.New("error getting session")
	}

	var isAuth bool
	if value, ok := session.Values["authenticated"].(bool); ok {
		isAuth = value
	} else {
		log.Println("error getting 'authenticated' value from session")
		return SessionInfo{}, errors.New("error getting 'authenticated' value from session")
	}

	var id int
	if value, ok := session.Values["user_id"].(int); ok {
		id = value
	} else {
		log.Println("error getting 'user_email' value from session")
		return SessionInfo{}, errors.New("error getting 'user_email' value from session")
	}

	var email string
	if value, ok := session.Values["user_email"].(string); ok {
		email = value
	} else {
		log.Println("error getting 'user_email' value from session")
		return SessionInfo{}, errors.New("error getting 'user_email' value from session")
	}

	var username string
	if value, ok := session.Values["username"].(string); ok {
		username = value
	} else {
		log.Println("error getting 'username' value from session")
		return SessionInfo{}, errors.New("error getting 'username' value from session")
	}

	s := SessionInfo{
		Authenticated: isAuth,
		Id:            id,
		Email:         email,
		Username:      username,
	}

	return s, nil
}
