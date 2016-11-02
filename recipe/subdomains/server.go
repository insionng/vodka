package main

import (
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	"github.com/insionng/vodka/middleware"
)

type (
	Host struct {
		Vodka *vodka.Vodka
	}
)

func main() {
	// Hosts
	hosts := make(map[string]*Host)

	//-----
	// API
	//-----

	api := vodka.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	hosts["api.localhost:1323"] = &Host{api}

	api.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "API")
	})

	//------
	// Blog
	//------

	blog := vodka.New()
	blog.Use(middleware.Logger())
	blog.Use(middleware.Recover())

	hosts["blog.localhost:1323"] = &Host{blog}

	blog.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "Blog")
	})

	//---------
	// Website
	//---------

	site := vodka.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())

	hosts["localhost:1323"] = &Host{site}

	site.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "Website")
	})

	// Server
	e := vodka.New()
	e.Any("/*", func(c vodka.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host()]

		if host == nil {
			err = vodka.ErrNotFound
		} else {
			host.Vodka.ServeHTTP(req, res)
		}

		return
	})
	e.Run(standard.New(":1323"))
}
