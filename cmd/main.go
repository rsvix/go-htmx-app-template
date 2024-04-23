package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rsvix/go-htmx-app-template/internal/handlers"
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

	app := echo.New()
	app.Debug = true
	app.Static("static", "./static")
	app.File("/favicon.ico", "./static/images/icon.ico")
	db := db.Connect()

	// Ip extractor (https://echo.labstack.com/docs/ip-address) - Not using, check /internal/utils/env_var.go
	// app.IPExtractor = echo.ExtractIPDirect()
	// app.IPExtractor = echo.ExtractIPFromXFFHeader()
	// app.IPExtractor = echo.ExtractIPFromRealIPHeader()

	// app.Pre(middleware.HTTPSRedirect())

	app.Use(
		middleware.Logger(),
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
		return c.Redirect(http.StatusSeeOther, "/notfound")
	}

	app.GET("/login", handlers.GetLoginHandler().Serve, middlewares.MustNotBeLogged())
	app.POST("/login", handlers.PostLoginHandler().Serve, middlewares.MustNotBeLogged())
	app.GET("/register", handlers.GetRegisterHandler().Serve, middlewares.MustNotBeLogged())
	app.POST("/register", handlers.PostRegisterHandler().Serve, middlewares.MustNotBeLogged())
	app.GET("/reset", handlers.GetResetHandler().Serve, middlewares.MustNotBeLogged())
	app.POST("/reset", handlers.PostResetHandler().Serve, middlewares.MustNotBeLogged())

	app.GET("/tkn/:token", handlers.GetTokenHandler().Serve)
	app.GET("/activate", handlers.GetActivateHandler().Serve)
	app.GET("/newactivation", handlers.GetNewActivationHandler().Serve)
	app.GET("/resetform", handlers.GetResetformHandler().Serve)
	app.POST("/resetform", handlers.PostResetformHandler().Serve)

	app.GET("/notfound", handlers.GetNotfoundHandler().Serve)
	app.GET("/error", handlers.GetInternalErrorHandler().Serve)
	app.GET("/terms", handlers.GetTermsHandlerParams().Serve)

	app.GET("/", handlers.GetIndexHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippets", handlers.GetSnippetsHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetform", handlers.GetSnippetFormHandler().Serve, middlewares.MustBeLogged())
	app.POST("/snippetform", handlers.PostSnippetFormHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetview/:id", handlers.GetSnippetViewHandler().Serve, middlewares.MustBeLogged())
	app.GET("/snippetedit/:id", handlers.GetSnippetEditHandler().Serve, middlewares.MustBeLogged())
	app.PUT("/snippetedit/:id", handlers.PutSnippetEditHandler().Serve, middlewares.MustBeLogged())
	app.GET("/account", handlers.GetAccountHandler().Serve, middlewares.MustBeLogged())
	app.GET("/edit_account", handlers.GetEditAccountHandler().Serve, middlewares.MustBeLogged())
	app.POST("/edit_account", handlers.PostEditAccountHandler().Serve, middlewares.MustBeLogged())
	app.GET("/logout", handlers.GetLogoutHandler().Serve, middlewares.MustBeLogged(), middlewares.NoCacheHeaders())

	log.Printf("Starting %v server on port %v", appName, appPort)
	app.Logger.Fatal(app.Start(":" + appPort))
}
