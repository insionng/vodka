package main

import (
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	"github.com/insionng/vodka/middleware"
)

var (
	users = []string{"Joe", "Veer", "Zion"}
)

func getUsers(c vodka.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {
	e := vodka.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// e.Use(middleware.CORS())

	// CORS restricted
	// Allows requests from any `https://insionng.com` or `https://insionng.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://insionng.com", "https://insionng.net"},
		AllowMethods: []string{vodka.GET, vodka.PUT, vodka.POST, vodka.DELETE},
	}))

	e.GET("/api/users", getUsers)
	e.Run(standard.New(":1323"))
}
