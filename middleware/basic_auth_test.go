package middleware

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	res := test.NewResponseRecorder()
	c := e.NewContext(req, res)
	f := func(u, p string) bool {
		if u == "joe" && p == "secret" {
			return true
		}
		return false
	}
	h := BasicAuth(f)(func(c vodka.Context) error {
		return c.String(http.StatusOK, "test")
	})

	// Valid credentials
	auth := basic + " " + base64.StdEncoding.EncodeToString([]byte("joe:secret"))
	req.Header().Set(vodka.HeaderAuthorization, auth)
	assert.NoError(t, h(c))

	// Incorrect password
	auth = basic + " " + base64.StdEncoding.EncodeToString([]byte("joe:password"))
	req.Header().Set(vodka.HeaderAuthorization, auth)
	he := h(c).(*vodka.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, he.Code)
	assert.Equal(t, basic+" realm=Restricted", res.Header().Get(vodka.HeaderWWWAuthenticate))

	// Empty Authorization header
	req.Header().Set(vodka.HeaderAuthorization, "")
	he = h(c).(*vodka.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, he.Code)

	// Invalid Authorization header
	auth = base64.StdEncoding.EncodeToString([]byte("invalid"))
	req.Header().Set(vodka.HeaderAuthorization, auth)
	he = h(c).(*vodka.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, he.Code)
}
