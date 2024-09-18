package main

import (
	"fmt"
	"net/http"

	"github.com/camdenwithrow/rdmapp/config"
	"github.com/camdenwithrow/rdmapp/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	hosts := map[string]*Host{}
	cfg := config.GetConfig()

	store := db.NewSQLStore()
	defer store.Close()
	store.GetUsers()

	admin := echo.New()
	admin.Use(middleware.Logger())
	admin.Use(middleware.Recover())
	hosts[fmt.Sprintf("admin.%s:%s", cfg.PublicHost, cfg.Port)] = &Host{admin}

	admin.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin")
	})

	site := echo.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())
	hosts[fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.Port)] = &Host{site}

	site.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Website")
	})

	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
