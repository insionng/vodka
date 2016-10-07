package middleware

import "github.com/insionng/vodka"

type (
	// Skipper defines a function to skip middleware. Returning true skips processing
	// the middleware.
	Skipper func(c vodka.Context) bool
)

// defaultSkipper returns false which processes the middleware.
func defaultSkipper(c vodka.Context) bool {
	return false
}
