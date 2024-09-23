package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/camdenwithrow/rdmapp/config"
	devDB "github.com/camdenwithrow/rdmapp/db/sqlite/dev"
	"github.com/camdenwithrow/rdmapp/handlers"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	hosts := map[string]*Host{}
	cfg := config.GetConfig()

	// store := dev.NewDevSQLiteStore()
	store := devDB.NewDevSQLiteStore()
	defer store.Close()
	// store.GetFeatures()
	// store.GetUsers()
	handler := handlers.New(store)

	admin := echo.New()
	admin.Use(middleware.Logger())
	admin.Use(middleware.Recover())
	hosts[fmt.Sprintf("admin.%s:%s", cfg.PublicHost, cfg.Port)] = &Host{admin}

	admin.GET("/", handler.AdminDashHandler)

	site := echo.New()
	// site.Use(middleware.Logger())
	// site.Use(middleware.Recover())
	hosts[fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.Port)] = &Host{site}

	site.GET("/:id", handler.RoadmapHandler)

	e := echo.New()
	e.Static("/static", "static")

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
