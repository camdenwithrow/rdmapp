package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
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
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "4321"
	}

	hostDomain, exists := os.LookupEnv("HOST")
	if !exists {
		hostDomain = "localhost"
	}

	admin := echo.New()
	admin.Use(middleware.Logger())
	admin.Use(middleware.Recover())
	hosts[fmt.Sprintf("admin.%s:%s", hostDomain, port)] = &Host{admin}

	admin.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin")
	})

	site := echo.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())
	hosts[fmt.Sprintf("%s:%s", hostDomain, port)] = &Host{site}

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

	e.Logger.Fatal(e.Start(":4321"))
}
