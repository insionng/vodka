package main

import (
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	"github.com/insionng/vodka/middleware"
)

func main() {
	// Vodka instance
	e := vodka.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Run(standard.New(":1323"))
}
