package main

import (
	"net/http"

	"github.com/insionng/vodka"
	mw "github.com/insionng/vodka/middleware"
)

// Handler
func hello(c *vodka.Context) error {
	return c.String(http.StatusOK, "Hello, Vodka!\n")
}

func main() {
	// Echo instance
	e := vodka.New()

	// Debug mode
	e.SetDebug(true)

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

	//-------
	// Slash
	//-------

	e.Use(mw.StripTrailingSlash())

	// or

	//	e.Use(mw.RedirectToSlash())

	// Gzip
	e.Use(mw.Gzip())

	// Routes
	e.Get("/", hello)

	// Start server
	e.Run(":1323")
}
