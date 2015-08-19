package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/insionng/vodka"
	mw "github.com/insionng/vodka/middleware"
)

const (
	Bearer     = "Bearer"
	SigningKey = "somethingsupersecret"
)

// A JSON Web Token middleware
func JWTAuth(key string) vodka.HandlerFunc {
	return func(c *vodka.Context) error {

		// Skip WebSocket
		if (c.Request().Header.Get(vodka.Upgrade)) == vodka.WebSocket {
			return nil
		}

		auth := c.Request().Header.Get("Authorization")
		l := len(Bearer)
		he := vodka.NewHTTPError(http.StatusUnauthorized)

		if len(auth) > l+1 && auth[:l] == Bearer {
			t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {

				// Always check the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// Return the key for validation
				return []byte(key), nil
			})
			if err == nil && t.Valid {
				// Store token claims in vodka.Context
				c.Set("claims", t.Claims)
				return nil
			}
		}
		return he
	}
}

func accessible(c *vodka.Context) error {
	return c.String(http.StatusOK, "No auth required for this route.\n")
}

func restricted(c *vodka.Context) error {
	return c.String(http.StatusOK, "Access granted with JWT.\n")
}

func main() {
	// Echo instance
	e := vodka.New()

	// Logger
	e.Use(mw.Logger())

	// Unauthenticated route
	e.Get("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(JWTAuth(SigningKey))
	r.Get("", restricted)

	// Start server
	e.Run(":1323")
}
