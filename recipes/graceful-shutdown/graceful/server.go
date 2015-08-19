package main

import (
	"net/http"
	"time"

	"github.com/insionng/vodka"
	"github.com/tylerb/graceful"
)

func main() {
	// Setup
	e := vodka.New()
	e.Get("/", func(c *vodka.Context) error {
		return c.String(http.StatusOK, "Sue sews rose on slow jor crows nose")
	})

	graceful.ListenAndServe(e.Server(":1323"), 5*time.Second)
}
