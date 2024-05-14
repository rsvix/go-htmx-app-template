package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rsvix/go-htmx-app-template/internal/handlers/accounthandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/activationhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/indexhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/internalerrorhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/ldaploginhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/loginhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/logouthandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/notfoundhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/registerhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/resethandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/snippetshandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/termshandler"
	"github.com/rsvix/go-htmx-app-template/internal/templates"

	"github.com/rsvix/go-htmx-app-template/internal/middlewares"
	"github.com/rsvix/go-htmx-app-template/internal/scheduler"
	"github.com/rsvix/go-htmx-app-template/internal/store/cookiestore"
	"github.com/rsvix/go-htmx-app-template/internal/store/db"
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
	appPort := utils.GetSetEnv("APP_PORT", "8080")
	appName := utils.GetSetEnv("APP_NAME", "GoBot")
	utils.GetSetEnv("POSTGRES_DB", "postgres")
	utils.GetSetEnv("POSTGRES_PORT", "5432")
	utils.GetSetEnv("POSTGRES_USER", "admin")
	utils.GetSetEnv("POSTGRES_PASSWORD", "123")
	utils.GetSetEnv("POSTGRES_HOST", "localhost")
	utils.GetSetEnv("APP_NAME_DB", "app-01")
	utils.GetSetEnv("DB_CONTEXT_KEY", "__db")

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
	db := db.ConnectMysql()

	app.HTTPErrorHandler = customHTTPErrorHandler

	// // This is for future stuff iam trying
	// var dbStats []struct {
	// 	Pid             uint
	// 	Datname         string
	// 	Usename         string
	// 	ApplicationName string
	// 	// ClientHostname  string
	// 	ClientPort   uint
	// 	BackendStart string
	// 	// QueryStart      string
	// 	// Query           string
	// 	State string
	// }
	// db.Raw("SELECT pid, datname, usename, application_name, client_port, backend_start, state FROM pg_stat_activity ORDER BY pid;").Scan(&dbStats)
	// // log.Printf("Apps connected: %v\n", dbStats)
	// for _, value := range dbApps {
	// 	if value.ApplicationName == "app-01" {
	// 		log.Printf("This is the main instance of the app, it will manage cronjobs\n")
	// 	}
	// }

	var dbStats []struct {
		Id      uint
		User    string
		Host    string
		Db      string
		Command string
		Time    string
		State   string
		Info    string
	}
	db.Raw("SELECT id, user, host, db, state FROM information_schema.processlist;").Scan(&dbStats)
	log.Printf("Apps connected: %v\n", dbStats)

	// https://github.com/go-co-op/gocron-gorm-lock

	sched := scheduler.BuildAsyncSched()

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

	echo.NotFoundHandler = func(c echo.Context) error {
		pageUrl := c.Request().URL
		c.Logger().Error("Page not found: %v\n", pageUrl)
		return c.Redirect(http.StatusSeeOther, "/notfound")
	}

	app.GET("/notfound", notfoundhandler.GetNotfoundHandler().Serve)
	app.GET("/error", internalerrorhandler.GetInternalErrorHandler().Serve)
	app.GET("/terms", termshandler.GetTermsHandlerParams().Serve)

	if ldapLoginBool {
		app.GET("/login", ldaploginhandler.GetLdapLoginHandler().Serve, middlewares.MustNotBeLogged())
		app.POST("/login", ldaploginhandler.PostLdapLoginHandler().Serve, middlewares.MustNotBeLogged())
	} else {
		app.GET("/login", loginhandler.GetLoginHandler().Serve, middlewares.MustNotBeLogged())
		app.POST("/login", loginhandler.PostLoginHandler().Serve, middlewares.MustNotBeLogged())
		app.GET("/register", registerhandler.GetRegisterHandler().Serve, middlewares.MustNotBeLogged())
		app.POST("/register", registerhandler.PostRegisterHandler().Serve, middlewares.MustNotBeLogged())
		app.GET("/reset", resethandler.GetResetHandler().Serve, middlewares.MustNotBeLogged())
		app.POST("/reset", resethandler.PostResetHandler().Serve, middlewares.MustNotBeLogged())
		app.GET("/resetform", resethandler.GetResetformHandler().Serve, middlewares.MustNotBeLogged())
		app.POST("/resetform", resethandler.PostResetformHandler().Serve, middlewares.MustNotBeLogged())
		app.GET("/pwreset", resethandler.GetResetTokenHandler().Serve, middlewares.MustNotBeLogged())
		app.GET("/activation", activationhandler.GetActivationTokenHandler().Serve, middlewares.MustNotBeLogged())
		app.GET("/newtoken", activationhandler.GetNewActivationHandler().Serve, middlewares.MustNotBeLogged())

		app.GET("/account", accounthandler.GetAccountHandler().Serve, middlewares.MustBeLogged())
		app.GET("/edit_account", accounthandler.GetEditAccountHandler().Serve, middlewares.MustBeLogged())
		app.POST("/edit_account", accounthandler.PostEditAccountHandler().Serve, middlewares.MustBeLogged())
		app.PUT("/edit_account", accounthandler.PutEditAccountHandler().Serve, middlewares.MustBeLogged())
	}

	app.GET("/", indexhandler.GetIndexHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippets", snippetshandler.GetSnippetsHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetform", snippetshandler.GetSnippetFormHandler().Serve, middlewares.MustBeLogged())
	app.POST("/snippetform", snippetshandler.PostSnippetFormHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetview/:id", snippetshandler.GetSnippetViewHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetedit/:id", snippetshandler.GetSnippetEditHandler().Serve, middlewares.MustBeLogged())
	app.PUT("/snippetedit/:id", snippetshandler.PutSnippetEditHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetdelete/:id", snippetshandler.GetSnippetDeleteModalEditHandler().Serve, middlewares.MustBeLogged())
	app.DELETE("/snippetdelete/:id", snippetshandler.DeleteSnippetEditHandler().Serve, middlewares.MustBeLogged())
	app.GET("/logout", logouthandler.GetLogoutHandler().Serve, middlewares.MustBeLogged())

	// log.Printf("Starting %v server on port %v", appName, appPort)
	// app.Logger.Fatal(app.Start(":" + appPort))

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
