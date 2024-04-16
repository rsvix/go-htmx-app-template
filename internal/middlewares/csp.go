package middlewares

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type key string

var NonceKey key = "nonces"

type Nonces struct {
	NavBar          string
	Htmx            string
	ResponseTargets string
	Tw              string
	Fa              string
	HtmxCSSHash     string
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func CSPMiddlewareCP(next http.Handler) http.Handler {
	// To use the same nonces in all responses, move the Nonces
	// struct creation to here, outside the handler.

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new Nonces struct for every request when here.
		// move to outside the handler to use the same nonces in all responses
		nonceSet := Nonces{
			Htmx:            generateRandomString(16),
			ResponseTargets: generateRandomString(16),
			Tw:              generateRandomString(16),
			Fa:              generateRandomString(16),
			HtmxCSSHash:     "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg=",
		}

		// set nonces in context
		ctx := context.WithValue(r.Context(), NonceKey, nonceSet)
		// insert the nonces into the content security policy header
		cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' 'nonce-%s' '%s';",
			nonceSet.Htmx,
			nonceSet.ResponseTargets,
			nonceSet.Tw,
			nonceSet.Fa,
			nonceSet.HtmxCSSHash)
		w.Header().Set("Content-Security-Policy", cspHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// https://stackoverflow.com/questions/69326129/does-set-method-of-echo-context-saves-the-value-to-the-underlying-context-cont
// https://rohitbels.medium.com/scrip-src-nonce-or-hash-algorithm-e43a6681f188

func CSPMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			nonceSet := Nonces{
				NavBar:          generateRandomString(16),
				Htmx:            generateRandomString(16),
				ResponseTargets: generateRandomString(16),
				Tw:              generateRandomString(16),
				Fa:              generateRandomString(16),
				HtmxCSSHash:     "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg=",
			}

			v := reflect.ValueOf(nonceSet)
			typeOfS := v.Type()
			for i := 0; i < v.NumField(); i++ {
				c.Set(typeOfS.Field(i).Name, v.Field(i).Interface())
			}

			cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' 'nonce-%s' '%s';",
				nonceSet.NavBar,
				nonceSet.Htmx,
				nonceSet.ResponseTargets,
				nonceSet.Tw,
				nonceSet.Fa,
				nonceSet.HtmxCSSHash)
			c.Response().Header().Set("Content-Security-Policy", cspHeader)

			return next(c)
		}
	}
}

func GetNavBarNonce(c echo.Context) string {
	nonce := c.Get("NavBar").(string)
	return nonce
}

func GetHtmxNonce(c echo.Context) string {
	nonce := c.Get("Htmx").(string)
	return nonce
}

func GetResponseTargetsNonce(c echo.Context) string {
	nonce := c.Get("ResponseTargets").(string)
	return nonce
}

func GetTwNonce(c echo.Context) string {
	nonce := c.Get("Tw").(string)
	return nonce
}

func GetFaNonce(c echo.Context) string {
	nonce := c.Get("Fa").(string)
	return nonce
}
