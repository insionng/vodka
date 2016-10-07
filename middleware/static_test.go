package middleware

import (
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestStatic(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := Static("../_fixture")(func(c vodka.Context) error {
		return vodka.ErrNotFound
	})

	// Directory
	if assert.NoError(t, h(c)) {
		assert.Contains(t, rec.Body.String(), "Vodka")
	}

	// HTML5 mode
	req = test.NewRequest(vodka.GET, "/client", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	static := StaticWithConfig(StaticConfig{
		Root:  "../_fixture",
		HTML5: true,
	})
	h = static(func(c vodka.Context) error {
		return vodka.ErrNotFound
	})
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Status())
	}

	// Browse
	req = test.NewRequest(vodka.GET, "/", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	static = StaticWithConfig(StaticConfig{
		Root:   "../_fixture/images",
		Browse: true,
	})
	h = static(func(c vodka.Context) error {
		return vodka.ErrNotFound
	})
	if assert.NoError(t, h(c)) {
		assert.Contains(t, rec.Body.String(), "walle")
	}

	// Not found
	req = test.NewRequest(vodka.GET, "/not-found", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	static = StaticWithConfig(StaticConfig{
		Root: "../_fixture/images",
	})
	h = static(func(c vodka.Context) error {
		return vodka.ErrNotFound
	})
	assert.Error(t, h(c))
}
