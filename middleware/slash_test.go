package middleware

import (
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestAddTrailingSlash(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/add-slash", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := AddTrailingSlash()(func(c vodka.Context) error {
		return nil
	})
	h(c)
	assert.Equal(t, "/add-slash/", req.URL().Path())
	assert.Equal(t, "/add-slash/", req.URI())

	// With config
	req = test.NewRequest(vodka.GET, "/add-slash?key=value", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = AddTrailingSlashWithConfig(TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	})(func(c vodka.Context) error {
		return nil
	})
	h(c)
	assert.Equal(t, http.StatusMovedPermanently, rec.Status())
	assert.Equal(t, "/add-slash/?key=value", rec.Header().Get(vodka.HeaderLocation))
}

func TestRemoveTrailingSlash(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/remove-slash/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := RemoveTrailingSlash()(func(c vodka.Context) error {
		return nil
	})
	h(c)
	assert.Equal(t, "/remove-slash", req.URL().Path())
	assert.Equal(t, "/remove-slash", req.URI())

	// With config
	req = test.NewRequest(vodka.GET, "/remove-slash/?key=value", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = RemoveTrailingSlashWithConfig(TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	})(func(c vodka.Context) error {
		return nil
	})
	h(c)
	assert.Equal(t, http.StatusMovedPermanently, rec.Status())
	assert.Equal(t, "/remove-slash?key=value", rec.Header().Get(vodka.HeaderLocation))

	// With bare URL
	req = test.NewRequest(vodka.GET, "http://localhost", nil)
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = RemoveTrailingSlash()(func(c vodka.Context) error {
		return nil
	})
	h(c)
	assert.Equal(t, "", req.URL().Path())
}
