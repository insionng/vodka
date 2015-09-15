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

	// Debug mode
	e.Debug()

	//------------
	// Middleware
	//------------

	// Logger
	e.Use(mw.Logger())

	// Recover
	e.Use(mw.Recover())

	// Basic auth
	e.Use(mw.BasicAuth(func(usr, pwd string) bool {
		if usr == "joe" && pwd == "secret" {
			return true
		}
		return false
	}))

	// Gzip
	e.Use(mw.Gzip())

	// Routes
	e.Get("/", hello)

	// Start server
	e.Run(":1323")
}
