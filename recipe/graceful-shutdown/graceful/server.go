package main

import (
	"net/http"
	"time"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	"github.com/tylerb/graceful"
)

func main() {
	// Setup
	e := vodka.New()
	e.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "Sue sews rose on slow joe crows nose")
	})
	std := standard.New(":1323")
	std.SetHandler(e)
	graceful.ListenAndServe(std.Server, 5*time.Second)
}
