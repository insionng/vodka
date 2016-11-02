package main

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
)

func main() {
	e := vodka.New()
	e.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "Six sick bricks tick")
	})
	std := standard.New(":1323")
	std.SetHandler(e)
	gracehttp.Serve(std.Server)
}
