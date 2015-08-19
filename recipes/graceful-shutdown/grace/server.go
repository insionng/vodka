package main

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/insionng/vodka"
)

func main() {
	// Setup
	e := vodka.New()
	e.Get("/", func(c *vodka.Context) error {
		return c.String(http.StatusOK, "Six sick bricks tick")
	})

	gracehttp.Serve(e.Server(":1323"))
}
