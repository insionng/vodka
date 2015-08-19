package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/libraries/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	e := vodka.New()
	e.SetDebug(true)
	req, _ := http.NewRequest(vodka.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := vodka.NewContext(req, vodka.NewResponse(rec), e)
	h := func(c *vodka.Context) error {
		panic("test")
	}
	Recover()(h)(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "panic recover")
}
