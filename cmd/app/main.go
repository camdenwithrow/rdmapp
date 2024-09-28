package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/camdenwithrow/rdmapp/config"
	devDB "github.com/camdenwithrow/rdmapp/db/sqlite/dev"
	"github.com/camdenwithrow/rdmapp/handlers"
	"github.com/camdenwithrow/rdmapp/services/auth"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	hosts := map[string]*Host{}
	cfg := config.GetConfig()

	store := devDB.NewDevSQLiteStore()
	defer store.Close()

	handler := handlers.New(store)

	// Initialize auth
	auth.Init()

	admin := echo.New()
	admin.Use(middleware.Logger())
	admin.Use(middleware.Recover())
	hosts[fmt.Sprintf("admin.%s:%s", cfg.PublicHost, cfg.Port)] = &Host{admin}

	// Add auth routes to admin
	admin.GET("/login", loginPageHandler)
	admin.GET("/auth/github", auth.LoginHandler)
	admin.GET("/auth/github/callback", auth.CallbackHandler)
	admin.GET("/logout", auth.LogoutHandler)

	// Protect admin routes with AuthMiddleware
	adminGroup := admin.Group("")
	adminGroup.Use(auth.AuthMiddleware)
	adminGroup.GET("/", handler.AdminDashHandler)
	// Add other admin routes here, they will all be protected

	site := echo.New()
	hosts[fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.Port)] = &Host{site}

	// Remove auth from site routes
	site.GET("/", handler.RoadmapHandler)

	e := echo.New()
	e.Static("/static", "static")

	e.Any("/*", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			return echo.ErrNotFound
		}
		host.Echo.ServeHTTP(res, req)
		return nil
	})

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}

func loginPageHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `
		<html>
			<body>
				<h1>Login</h1>
				<a href="/auth/github">Login with GitHub</a>
			</body>
		</html>
	`)
}
