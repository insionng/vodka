package middleware

import "github.com/insionng/vodka"

// StripTrailingSlash returns a middleware which removes trailing slash from request
// path.
func StripTrailingSlash() vodka.HandlerFunc {
	return func(c *vodka.Context) error {
		p := c.Request().URL.Path
		l := len(p)
		if p[l-1] == '/' {
			c.Request().URL.Path = p[:l-1]
		}
		return nil
	}
}
