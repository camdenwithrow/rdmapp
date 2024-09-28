package auth

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"

	"github.com/camdenwithrow/rdmapp/config"
)

var (
	store *sessions.CookieStore
)

func Init() {
	cfg := config.GetConfig()

	store = sessions.NewCookieStore([]byte(cfg.SessionSecret))
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = cfg.Environment != "development" // Set to true if not in development

	gothic.Store = store

	gob.Register(goth.User{})

	callbackURL := fmt.Sprintf("http://admin.%s:%s/auth/github/callback", cfg.PublicHost, cfg.Port)
	goth.UseProviders(
		github.New(cfg.GithubClientID, cfg.GithubClientSecret, callbackURL),
	)
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "auth-session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}

func LoginHandler(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}

func CallbackHandler(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	session, _ := store.Get(c.Request(), "auth-session")
	session.Values["authenticated"] = true
	session.Values["user"] = user
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/")
}

func LogoutHandler(c echo.Context) error {
	session, _ := store.Get(c.Request(), "auth-session")
	session.Values["authenticated"] = false
	session.Values["user"] = nil
	session.Save(c.Request(), c.Response())

	gothic.Logout(c.Response(), c.Request())
	return c.Redirect(http.StatusSeeOther, "/")
}

func GetUser(c echo.Context) (goth.User, bool) {
	session, _ := store.Get(c.Request(), "auth-session")
	user, ok := session.Values["user"].(goth.User)
	return user, ok
}
