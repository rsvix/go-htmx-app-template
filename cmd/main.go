package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rsvix/go-htmx-app-template/cmd/routes"
	"github.com/rsvix/go-htmx-app-template/internal/middlewares"
	"github.com/rsvix/go-htmx-app-template/internal/scheduler"
	"github.com/rsvix/go-htmx-app-template/internal/store/cookiestore"
	"github.com/rsvix/go-htmx-app-template/internal/store/db"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	// If you want to get the error code and msg
	msg := ""
	code := http.StatusInternalServerError
	if e, ok := err.(*echo.HTTPError); ok {
		msg = e.Message.(string)
		code = e.Code
	}
	log.Printf("Echo handler error\ncode: %v\nmessage: %v\n", code, msg)

	c.Logger().Error(err)
	templates.ErrorPage(c, "Error", "We are working to fix the problem").Render(c.Request().Context(), c.Response())
}

func main() {
	instanceId := uuid.New().String()
	log.Printf("\n\nApp instance id: %v\n\n", instanceId)

	appPort := utils.GetSetEnv("APP_PORT", "8080")
	appName := utils.GetSetEnv("APP_NAME", "GoBot")
	utils.GetSetEnv("DB_CONTEXT_KEY", "__db")
	utils.GetSetEnv("IS_SCHEDULED", "false")

	utils.GetSetEnv("DB_URL", "mysql://admin:password123@tcp(localhost:3306)/testbg")

	utils.GetSetEnv("LDAP_URL", "ldap://localhost:389")
	utils.GetSetEnv("LDAP_BASE_DN", "DC=example,DC=com")
	utils.GetSetEnv("LDAP_GROUP", "OU=group,DC=example,DC=com")
	ldapLogin := utils.GetSetEnv("LDAP_LOGIN", "false")

	ldapLoginBool, err := strconv.ParseBool(ldapLogin)
	if err != nil {
		log.Println(err)
		ldapLoginBool = false
	}

	app := echo.New()
	// app.Debug = true
	app.Static("static", "./static")
	app.File("/favicon.ico", "./static/images/icon.ico")

	db := db.Connect()
	sched := scheduler.BuildAsyncSched(db, instanceId)

	app.HTTPErrorHandler = customHTTPErrorHandler

	// app.Pre(middleware.HTTPSRedirect())

	app.Use(
		middleware.Logger(),
		middleware.Recover(),
		middlewares.DatabaseMiddleware(db),
		session.Middleware(cookiestore.Start(db)),
		middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(15)),
		middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:      "1; mode=block",
			ContentTypeNosniff: "nosniff",
			XFrameOptions:      "",
			HSTSMaxAge:         3600,
			// ContentSecurityPolicy: "default-src 'self'", // defining in middlewares.CSPMiddleware()
		}),
		middlewares.CSPMiddleware(),
		middleware.CSRFWithConfig(middleware.CSRFConfig{
			// TokenLookup: "form:_csrf",
			TokenLookup:    "cookie:_csrf",
			CookiePath:     "/",
			CookieDomain:   "localhost",
			CookieMaxAge:   84600,
			CookieSecure:   false,
			CookieHTTPOnly: true,
			CookieSameSite: http.SameSiteStrictMode,
		}),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Skipper:      middleware.DefaultSkipper,
			ErrorMessage: "timeout error - unable to connect",
			OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
				log.Printf("Timeout error on path: %v\n", c.Path())
			},
			Timeout: 30 * time.Second,
		}),
	)

	// app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"*"},
	// }))

	app = routes.AppRoutes(app, ldapLoginBool)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		log.Printf("Starting %v server on port %v", appName, appPort)
		if err := app.Start(":" + appPort); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := sched.Shutdown(); err != nil {
		log.Println(err)
	}
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
