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
	// Not using - Cehck github.com/rsvix/go-htmx-app-template/internal/utils/env_var.go
	// app.IPExtractor = echo.ExtractIPDirect()
	// app.IPExtractor = echo.ExtractIPFromXFFHeader()
	// app.IPExtractor = echo.ExtractIPFromRealIPHeader()

	// Middlewares
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middlewares.DatabaseMiddleware(db))
	app.Use(session.Middleware(cookiestore.Start(db)))

	// app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	// Enable in production
	// app.Use(middleware.Secure())
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
	app.POST("/reset", handlers.NewPostResetHandler(handlers.PostResetHandlerParams{}).ServeHTTP)
	app.GET("/resetform", handlers.GetResetformHandler().Serve)
	app.POST("/resetform", handlers.NewPostProcessResetHandler(handlers.PostProcessResetHandlerParams{}).ServeHTTP)
	app.GET("/login", handlers.GetLoginHandler().Serve)
	app.POST("/login", handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{}).ServeHTTP)
	app.GET("/logout", handlers.GetLogoutHandler().Serve)
	app.GET("/tkn/:token", handlers.GetTokenHandler{}.ServeHTTP)

	echo.NotFoundHandler = func(c echo.Context) error {
		return templates.NotfoundPage("Not Found", "Page not found").Render(c.Request().Context(), c.Response())
	}

	log.Printf("Starting %v server on port %v", appName, appPort)
	app.Logger.Fatal(app.Start(":" + appPort))
}
