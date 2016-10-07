package middleware

import (
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestHTTPSRedirect(t *testing.T) {
	e := vodka.New()
	next := func(c vodka.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req := test.NewRequest(vodka.GET, "http://insionng.com", nil)
	res := test.NewResponseRecorder()
	c := e.NewContext(req, res)
	HTTPSRedirect()(next)(c)
	assert.Equal(t, http.StatusMovedPermanently, res.Status())
	assert.Equal(t, "https://insionng.com", res.Header().Get(vodka.HeaderLocation))
}

func TestHTTPSWWWRedirect(t *testing.T) {
	e := vodka.New()
	next := func(c vodka.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req := test.NewRequest(vodka.GET, "http://insionng.com", nil)
	res := test.NewResponseRecorder()
	c := e.NewContext(req, res)
	HTTPSWWWRedirect()(next)(c)
	assert.Equal(t, http.StatusMovedPermanently, res.Status())
	assert.Equal(t, "https://www.insionng.com", res.Header().Get(vodka.HeaderLocation))
}

func TestWWWRedirect(t *testing.T) {
	e := vodka.New()
	next := func(c vodka.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req := test.NewRequest(vodka.GET, "http://insionng.com", nil)
	res := test.NewResponseRecorder()
	c := e.NewContext(req, res)
	WWWRedirect()(next)(c)
	assert.Equal(t, http.StatusMovedPermanently, res.Status())
	assert.Equal(t, "http://www.insionng.com", res.Header().Get(vodka.HeaderLocation))
}

func TestNonWWWRedirect(t *testing.T) {
	e := vodka.New()
	next := func(c vodka.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req := test.NewRequest(vodka.GET, "http://www.insionng.com", nil)
	res := test.NewResponseRecorder()
	c := e.NewContext(req, res)
	NonWWWRedirect()(next)(c)
	assert.Equal(t, http.StatusMovedPermanently, res.Status())
	assert.Equal(t, "http://insionng.com", res.Header().Get(vodka.HeaderLocation))
}
