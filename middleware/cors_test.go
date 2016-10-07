package middleware

import (
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	cors := CORSWithConfig(CORSConfig{
		AllowCredentials: true,
	})
	h := cors(func(c vodka.Context) error {
		return c.String(http.StatusOK, "test")
	})

	// No origin header
	h(c)
	assert.Equal(t, "", rec.Header().Get(vodka.HeaderAccessControlAllowOrigin))

	// Empty origin header
	req = test.NewRequest(vodka.GET, "/", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	req.Header().Set(vodka.HeaderOrigin, "")
	h(c)
	assert.Equal(t, "*", rec.Header().Get(vodka.HeaderAccessControlAllowOrigin))

	// Wildcard origin
	req = test.NewRequest(vodka.GET, "/", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	req.Header().Set(vodka.HeaderOrigin, "localhost")
	h(c)
	assert.Equal(t, "*", rec.Header().Get(vodka.HeaderAccessControlAllowOrigin))

	// Simple request
	req = test.NewRequest(vodka.GET, "/", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	req.Header().Set(vodka.HeaderOrigin, "localhost")
	cors = CORSWithConfig(CORSConfig{
		AllowOrigins:     []string{"localhost"},
		AllowCredentials: true,
		MaxAge:           3600,
	})
	h = cors(func(c vodka.Context) error {
		return c.String(http.StatusOK, "test")
	})
	h(c)
	assert.Equal(t, "localhost", rec.Header().Get(vodka.HeaderAccessControlAllowOrigin))

	// Preflight request
	req = test.NewRequest(vodka.OPTIONS, "/", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	req.Header().Set(vodka.HeaderOrigin, "localhost")
	req.Header().Set(vodka.HeaderContentType, vodka.MIMEApplicationJSON)
	h(c)
	assert.Equal(t, "localhost", rec.Header().Get(vodka.HeaderAccessControlAllowOrigin))
	assert.NotEmpty(t, rec.Header().Get(vodka.HeaderAccessControlAllowMethods))
	assert.Equal(t, "true", rec.Header().Get(vodka.HeaderAccessControlAllowCredentials))
	assert.Equal(t, "3600", rec.Header().Get(vodka.HeaderAccessControlMaxAge))
}
