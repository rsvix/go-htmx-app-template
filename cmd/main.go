package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rsvix/go-htmx-app-template/internal/handlers/accounthandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/activationhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/indexhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/internalerrorhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/loginhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/logouthandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/notfoundhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/registerhandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/resethandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/snippetshandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/termshandler"
	"github.com/rsvix/go-htmx-app-template/internal/handlers/tokenhandler"

	"github.com/rsvix/go-htmx-app-template/internal/middlewares"
	"github.com/rsvix/go-htmx-app-template/internal/store/cookiestore"
	"github.com/rsvix/go-htmx-app-template/internal/store/db"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
)

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

	app := echo.New()
	// app.Debug = true
	app.Static("static", "./static")
	app.File("/favicon.ico", "./static/images/icon.ico")
	db := db.Connect()

	var dbApps []struct {
		Pid             uint
		Datname         string
		Usename         string
		ApplicationName string
		// ClientHostname  string
		ClientPort   uint
		BackendStart string
		// QueryStart      string
		// Query           string
		State string
	}
	db.Raw("SELECT pid, datname, usename, application_name, client_port, backend_start, state FROM pg_stat_activity ORDER BY pid;").Scan(&dbApps)
	// https://stackoverflow.com/questions/27435839/how-to-list-active-connections-on-postgresql
	// log.Printf("Apps connected: %v\n", dbApps)
	for _, value := range dbApps {
		if value.ApplicationName == "app-01" {
			log.Printf("This is the main instance of the app, it will manage cronjobs\n")
		}
	}

	// Ip extractor (https://echo.labstack.com/docs/ip-address) - Not using, check /internal/utils/env_var.go
	// app.IPExtractor = echo.ExtractIPDirect()
	// app.IPExtractor = echo.ExtractIPFromXFFHeader()
	// app.IPExtractor = echo.ExtractIPFromRealIPHeader()

	// app.Pre(middleware.HTTPSRedirect())

	app.Use(
		// middleware.Logger(),
		middleware.Recover(),
		middlewares.DatabaseMiddleware(db),
		session.Middleware(cookiestore.Start(db)),
		middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:      "1; mode=block",
			ContentTypeNosniff: "nosniff",
			XFrameOptions:      "",
			HSTSMaxAge:         3600,
			// ContentSecurityPolicy: "default-src 'self'",
		}),
		middlewares.CSPMiddleware(),
	)

	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		// TokenLookup: "form:_csrf",
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieDomain:   "localhost",
		CookieMaxAge:   84600,
		CookieSecure:   false,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	// app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"*"},
	// }))

	echo.NotFoundHandler = func(c echo.Context) error {
		// return c.Redirect(http.StatusSeeOther, "/notfound")
		return c.Redirect(http.StatusSeeOther, "/")
	}

	app.GET("/login", loginhandler.GetLoginHandler().Serve, middlewares.MustNotBeLogged())
	app.POST("/login", loginhandler.PostLoginHandler().Serve, middlewares.MustNotBeLogged())

	app.GET("/register", registerhandler.GetRegisterHandler().Serve, middlewares.MustNotBeLogged())
	app.POST("/register", registerhandler.PostRegisterHandler().Serve, middlewares.MustNotBeLogged())

	app.GET("/reset", resethandler.GetResetHandler().Serve, middlewares.MustNotBeLogged())
	app.POST("/reset", resethandler.PostResetHandler().Serve, middlewares.MustNotBeLogged())
	app.GET("/resetform", resethandler.GetResetformHandler().Serve)
	app.POST("/resetform", resethandler.PostResetformHandler().Serve)

	app.GET("/tkn/:token", tokenhandler.GetTokenHandler().Serve)

	app.GET("/activate", activationhandler.GetActivateHandler().Serve, middlewares.MustNotBeLogged())
	app.GET("/newactivation", activationhandler.GetNewActivationHandler().Serve, middlewares.MustNotBeLogged())
	app.GET("/actvtkn", activationhandler.GetActivationTokenHandler().Serve, middlewares.MustNotBeLogged())
	// app.GET("/activationtkn/:token", activationhandler.GetActivationTokenHandler().Serve, middlewares.MustNotBeLogged())

	app.GET("/notfound", notfoundhandler.GetNotfoundHandler().Serve)
	app.GET("/error", internalerrorhandler.GetInternalErrorHandler().Serve)

	app.GET("/terms", termshandler.GetTermsHandlerParams().Serve)

	app.GET("/", indexhandler.GetIndexHandler().Serve, middlewares.MustBeLogged())

	app.GET("/snippets", snippetshandler.GetSnippetsHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetform", snippetshandler.GetSnippetFormHandler().Serve, middlewares.MustBeLogged())
	app.POST("/snippetform", snippetshandler.PostSnippetFormHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetview/:id", snippetshandler.GetSnippetViewHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetedit/:id", snippetshandler.GetSnippetEditHandler().Serve, middlewares.MustBeLogged())
	app.PUT("/snippetedit/:id", snippetshandler.PutSnippetEditHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetdelete/:id", snippetshandler.GetSnippetDeleteModalEditHandler().Serve, middlewares.MustBeLogged())
	app.DELETE("/snippetdelete/:id", snippetshandler.DeleteSnippetEditHandler().Serve, middlewares.MustBeLogged())

	app.GET("/account", accounthandler.GetAccountHandler().Serve, middlewares.MustBeLogged())
	app.GET("/edit_account", accounthandler.GetEditAccountHandler().Serve, middlewares.MustBeLogged())
	app.POST("/edit_account", accounthandler.PostEditAccountHandler().Serve, middlewares.MustBeLogged())

	app.GET("/logout", logouthandler.GetLogoutHandler().Serve, middlewares.MustBeLogged())

	log.Printf("Starting %v server on port %v", appName, appPort)
	app.Logger.Fatal(app.Start(":" + appPort))
}
