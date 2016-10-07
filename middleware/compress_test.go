package middleware

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestGzip(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)

	// Skip if no Accept-Encoding header
	h := Gzip()(func(c vodka.Context) error {
		c.Response().Write([]byte("test")) // For Content-Type sniffing
		return nil
	})
	h(c)
	assert.Equal(t, "test", rec.Body.String())

	req = test.NewRequest(vodka.GET, "/", nil)
	req.Header().Set(vodka.HeaderAcceptEncoding, "gzip")
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)

	// Gzip
	h(c)
	assert.Equal(t, "gzip", rec.Header().Get(vodka.HeaderContentEncoding))
	assert.Contains(t, rec.Header().Get(vodka.HeaderContentType), vodka.MIMETextPlain)
	r, err := gzip.NewReader(rec.Body)
	defer r.Close()
	if assert.NoError(t, err) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		assert.Equal(t, "test", buf.String())
	}
}

func TestGzipNoContent(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := Gzip()(func(c vodka.Context) error {
		return c.NoContent(http.StatusOK)
	})
	if assert.NoError(t, h(c)) {
		assert.Empty(t, rec.Header().Get(vodka.HeaderContentEncoding))
		assert.Empty(t, rec.Header().Get(vodka.HeaderContentType))
		assert.Equal(t, 0, len(rec.Body.Bytes()))
	}
}

func TestGzipErrorReturned(t *testing.T) {
	e := vodka.New()
	e.Use(Gzip())
	e.GET("/", func(c vodka.Context) error {
		return vodka.NewHTTPError(http.StatusInternalServerError, "error")
	})
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	e.ServeHTTP(req, rec)
	assert.Empty(t, rec.Header().Get(vodka.HeaderContentEncoding))
	assert.Equal(t, "error", rec.Body.String())
}
