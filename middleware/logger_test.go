package middleware

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	// Note: Just for the test coverage, not a real test.
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := Logger()(func(c vodka.Context) error {
		return c.String(http.StatusOK, "test")
	})

	// Status 2xx
	h(c)

	// Status 3xx
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = Logger()(func(c vodka.Context) error {
		return c.String(http.StatusTemporaryRedirect, "test")
	})
	h(c)

	// Status 4xx
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = Logger()(func(c vodka.Context) error {
		return c.String(http.StatusNotFound, "test")
	})
	h(c)

	// Status 5xx with empty path
	req = test.NewRequest(vodka.GET, "", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = Logger()(func(c vodka.Context) error {
		return errors.New("error")
	})
	h(c)
}

func TestLoggerIPAddress(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	buf := new(bytes.Buffer)
	e.Logger().SetOutput(buf)
	ip := "127.0.0.1"
	h := Logger()(func(c vodka.Context) error {
		return c.String(http.StatusOK, "test")
	})

	// With X-Real-IP
	req.Header().Add(vodka.HeaderXRealIP, ip)
	h(c)
	assert.Contains(t, ip, buf.String())

	// With X-Forwarded-For
	buf.Reset()
	req.Header().Del(vodka.HeaderXRealIP)
	req.Header().Add(vodka.HeaderXForwardedFor, ip)
	h(c)
	assert.Contains(t, ip, buf.String())

	buf.Reset()
	h(c)
	assert.Contains(t, ip, buf.String())
}
