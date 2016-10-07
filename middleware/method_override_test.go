package middleware

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestMethodOverride(t *testing.T) {
	e := vodka.New()
	m := MethodOverride()
	h := func(c vodka.Context) error {
		return c.String(http.StatusOK, "test")
	}

	// Override with http header
	req := test.NewRequest(vodka.POST, "/", nil)
	rec := test.NewResponseRecorder()
	req.Header().Set(vodka.HeaderXHTTPMethodOverride, vodka.DELETE)
	c := e.NewContext(req, rec)
	m(h)(c)
	assert.Equal(t, vodka.DELETE, req.Method())

	// Override with form parameter
	m = MethodOverrideWithConfig(MethodOverrideConfig{Getter: MethodFromForm("_method")})
	req = test.NewRequest(vodka.POST, "/", bytes.NewReader([]byte("_method="+vodka.DELETE)))
	rec = test.NewResponseRecorder()
	req.Header().Set(vodka.HeaderContentType, vodka.MIMEApplicationForm)
	c = e.NewContext(req, rec)
	m(h)(c)
	assert.Equal(t, vodka.DELETE, req.Method())

	// Override with query paramter
	m = MethodOverrideWithConfig(MethodOverrideConfig{Getter: MethodFromQuery("_method")})
	req = test.NewRequest(vodka.POST, "/?_method="+vodka.DELETE, nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	m(h)(c)
	assert.Equal(t, vodka.DELETE, req.Method())

	// Ignore `GET`
	req = test.NewRequest(vodka.GET, "/", nil)
	req.Header().Set(vodka.HeaderXHTTPMethodOverride, vodka.DELETE)
	assert.Equal(t, vodka.GET, req.Method())
}
