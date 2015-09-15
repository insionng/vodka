package middleware

import (
	"encoding/base64"
	"net/http"

	"github.com/insionng/vodka"
)

type (
	BasicValidateFunc func(string, string) bool
)

const (
	Basic = "Basic"
)

// BasicAuth returns an HTTP basic authentication middleware.
//
// For valid credentials it calls the next handler.
// For invalid credentials, it sends "401 - Unauthorized" response.
func BasicAuth(fn BasicValidateFunc) vodka.HandlerFunc {
	return func(c *vodka.Context) error {
		// Skip WebSocket
		if (c.Request().Header.Get(vodka.Upgrade)) == vodka.WebSocket {
			return nil
		}

		auth := c.Request().Header.Get(vodka.Authorization)
		l := len(Basic)

		if len(auth) > l+1 && auth[:l] == Basic {
			b, err := base64.StdEncoding.DecodeString(auth[l+1:])
			if err == nil {
				cred := string(b)
				for i := 0; i < len(cred); i++ {
					if cred[i] == ':' {
						// Verify credentials
						if fn(cred[:i], cred[i+1:]) {
							return nil
						}
					}
				}
			}
		}
		c.Response().Header().Set(vodka.WWWAuthenticate, Basic+" realm=Restricted")
		return vodka.NewHTTPError(http.StatusUnauthorized)
	}
}
