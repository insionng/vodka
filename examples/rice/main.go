package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/insionng/vodka"
)

func main() {
	handler := http.StripPrefix(
		"/static/", http.FileServer(rice.MustFindBox("app").HTTPBox()),
	)
	e := vodka.New()
	e.Get("/static/*", func(c *vodka.Context) error {
		handler.ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})
	e.Run("localhost:1234")
}
