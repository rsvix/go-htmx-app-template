package middlewares

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	htmxCssHash = "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg="
)

// https://stackoverflow.com/questions/69326129/does-set-method-of-echo-context-saves-the-value-to-the-underlying-context-cont
// https://rohitbels.medium.com/scrip-src-nonce-or-hash-algorithm-e43a6681f188
// https://stackoverflow.com/questions/76270173/can-a-nonce-be-used-for-multiple-scripts-or-not

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func CSPMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			randomNonce := generateRandomString(16)

			// Storing nonce in the context
			// c.Set("randomNonce", randomNonce)

			// Storing nonce in session cookie
			session, err := session.Get("authenticate-sessions", c)
			if err != nil {
				log.Printf("Error getting session: %v\n", err)
				return err
			}
			session.Values["randomNonce"] = randomNonce
			if err := session.Save(c.Request(), c.Response()); err != nil {
				log.Printf("Error saving session: %s", err)
				return err
			}

			cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s'; style-src 'nonce-%s' '%s'; img-src '%s';",
				randomNonce,
				randomNonce,
				htmxCssHash,
				// randomNonce,
				"self",
			)
			c.Response().Header().Set("Content-Security-Policy", cspHeader)
			return next(c)
		}
	}
}

func GetRandomNonce(c echo.Context) string {
	// Getting nonce from context
	// nonce := c.Get("randomNonce").(string)
	// return nonce

	// Getting nonce from session cookie
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		return ""
	}
	if value, ok := session.Values["randomNonce"].(string); ok {
		return value
	}
	return ""
}
