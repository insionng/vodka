package middleware

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	e := vodka.New()
	buf := new(bytes.Buffer)
	e.SetLogOutput(buf)
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := Recover()(vodka.HandlerFunc(func(c vodka.Context) error {
		panic("test")
	}))
	h(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Status())
	assert.Contains(t, buf.String(), "PANIC RECOVER")
}
