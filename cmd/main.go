package main

import (
	"log"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rsvix/go-htmx-app-template/internal/handlers"
	"github.com/rsvix/go-htmx-app-template/internal/middlewares"
	"github.com/rsvix/go-htmx-app-template/internal/store/cookiestore"
	"github.com/rsvix/go-htmx-app-template/internal/store/db"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
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
	app.Static("static", "./static")
	app.File("/favicon.ico", "./static/images/icon.ico")
	db := db.Connect()

	// Ip extractor - https://echo.labstack.com/docs/ip-address
	// Not using - Check github.com/rsvix/go-htmx-app-template/internal/utils/env_var.go
	// app.IPExtractor = echo.ExtractIPDirect()
	// app.IPExtractor = echo.ExtractIPFromXFFHeader()
	// app.IPExtractor = echo.ExtractIPFromRealIPHeader()

	// Middlewares
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middlewares.DatabaseMiddleware(db))
	app.Use(session.Middleware(cookiestore.Start(db)))
	// app.Use(middlewares.TextHTMLMiddleware())
	// app.Use(middlewares.CSPMiddleware())

	// app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	app.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'; script-src 'nonce-2726c7f26c'; style-src 'nonce-2726c7f26s' 'sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg=';",
	}))

	// app.Pre(middleware.HTTPSRedirect())

	// Allow CORS For testing - Comment this in production
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	// Handlers
	app.GET("/", handlers.GetIndexHandler().Serve)
	app.GET("/register", handlers.GetRegisterHandler().Serve)
	app.POST("/register", handlers.PostRegisterHandler().Serve)
	app.GET("/activate", handlers.GetActivateHandler().Serve)
	app.GET("/newactivation", handlers.GetNewActivationHandler().Serve)
	app.GET("/reset", handlers.GetResetHandler().Serve)
	app.POST("/reset", handlers.PostResetHandler().Serve)
	app.GET("/resetform", handlers.GetResetformHandler().Serve)
	app.POST("/resetform", handlers.PostResetformHandler().Serve)
	app.GET("/login", handlers.GetLoginHandler().Serve)
	app.POST("/login", handlers.PostLoginHandler().Serve)
	app.GET("/logout", handlers.GetLogoutHandler().Serve, middlewares.NoCache())
	app.GET("/tkn/:token", handlers.GetTokenHandler().Serve)

	// noCacheGroup := app.Group(func(c echo.Context), middlewares.NoCache())

	echo.NotFoundHandler = func(c echo.Context) error {
		return templates.NotfoundPage("Not Found", "Page not found").Render(c.Request().Context(), c.Response())
	}

	log.Printf("Starting %v server on port %v", appName, appPort)
	app.Logger.Fatal(app.Start(":" + appPort))
}
