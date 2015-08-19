package main

import (
	"net/http"

	"github.com/insionng/vodka"
	mw "github.com/insionng/vodka/middleware"
)

// Handler
func hello(c *vodka.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	// Echo instance
	e := vodka.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Get("/", hello)

	// Start server
	e.Run(":1323")
}
