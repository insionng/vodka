package main

import (
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/middleware"
)

type (
	user struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users map[string]user
)

func init() {
	users = map[string]user{
		"1": user{
			ID:   "1",
			Name: "Wreck-It Ralph",
		},
	}

	// hook into the vodka instance to create an endpoint group
	// and add specific middleware to it plus handlers
	g := e.Group("/users")
	g.Use(middleware.CORS())

	g.POST("", createUser)
	g.GET("", getUsers)
	g.GET("/:id", getUser)
}

func createUser(c vodka.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = *u
	return c.JSON(http.StatusCreated, u)
}

func getUsers(c vodka.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUser(c vodka.Context) error {
	return c.JSON(http.StatusOK, users[c.P(0)])
}
