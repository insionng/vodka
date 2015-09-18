package middleware

import (
	"log"
	"net"
	"time"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/libraries/labstack/gommon/color"
)

func Logger() vodka.MiddlewareFunc {
	return func(h vodka.HandlerFunc) vodka.HandlerFunc {
		return func(c *vodka.Context) error {

			remoteAddr := c.Request().RemoteAddr
			if realIP := c.Request().Header.Get("X-Real-IP"); realIP != "" {
				remoteAddr = realIP
			}
			if realIP := c.Request().Header.Get("X-Forwarded-For"); realIP != "" {
				remoteAddr = realIP
			}

			start := time.Now()
			if err := h(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			method := c.Request().Method
			path := c.Request().URL.Path
			if path == "" {
				path = "/"
			}
			size := c.Response().Size()

			n := c.Response().Status()
			code := color.Green(n)
			switch {
			case n >= 500:
				code = color.Red(n)
			case n >= 400:
				code = color.Yellow(n)
			case n >= 300:
				code = color.Cyan(n)
			}

			remoteAddr, _, _ = net.SplitHostPort(remoteAddr)

			log.Printf("%s %s %s %s %d", method, path, code, stop.Sub(start), size)
			return nil
		}
	}
}
