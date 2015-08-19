package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/insionng/testify/assert"
	"github.com/insionng/vodka"
)

func TestStripTrailingSlash(t *testing.T) {
	req, _ := http.NewRequest(vodka.GET, "/users/", nil)
	rec := httptest.NewRecorder()
	c := vodka.NewContext(req, vodka.NewResponse(rec), vodka.New())
	StripTrailingSlash()(c)
	assert.Equal(t, "/users", c.Request().URL.Path)
}
