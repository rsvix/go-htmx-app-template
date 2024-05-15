package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	"github.com/rsvix/go-htmx-app-template/internal/middlewares"
)

func AppRoutes(app *echo.Echo, ldapLoginBool bool) *echo.Echo {
	echo.NotFoundHandler = func(c echo.Context) error {
		pageUrl := c.Request().URL
		c.Logger().Error("Page not found: %v\n", pageUrl)
		return c.Redirect(http.StatusSeeOther, "/notfound")
	}

	// Open routes
	app.GET("/notfound", notfoundhandler.GetNotfoundHandler().Serve)
	app.GET("/error", internalerrorhandler.GetInternalErrorHandler().Serve)
	app.GET("/terms", termshandler.GetTermsHandlerParams().Serve)

	if ldapLoginBool {
		// Not logged in routes
		app.GET("/login", ldaploginhandler.GetLdapLoginHandler().Serve, middlewares.MustNotBeLogged())
		app.POST("/login", ldaploginhandler.PostLdapLoginHandler().Serve, middlewares.MustNotBeLogged())
	} else {
		// Not logged in routes
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
		// Logged in routes
		app.GET("/account", accounthandler.GetAccountHandler().Serve, middlewares.MustBeLogged())
		app.GET("/edit_account", accounthandler.GetEditAccountHandler().Serve, middlewares.MustBeLogged())
		app.POST("/edit_account", accounthandler.PostEditAccountHandler().Serve, middlewares.MustBeLogged())
		app.PUT("/edit_account", accounthandler.PutEditAccountHandler().Serve, middlewares.MustBeLogged())
	}
	// Logged in routes
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

	return app
}
