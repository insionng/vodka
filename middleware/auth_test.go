package middleware

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/libraries/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {
	req, _ := http.NewRequest(vodka.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := vodka.NewContext(req, vodka.NewResponse(rec), vodka.New())
	fn := func(u, p string) bool {
		if u == "joe" && p == "secret" {
			return true
		}
		return false
	}
	ba := BasicAuth(fn)

	// Valid credentials
	auth := Basic + " " + base64.StdEncoding.EncodeToString([]byte("joe:secret"))
	req.Header.Set(vodka.Authorization, auth)
	assert.NoError(t, ba(c))

	//---------------------
	// Invalid credentials
	//---------------------

	// Incorrect password
	auth = Basic + " " + base64.StdEncoding.EncodeToString([]byte("joe:password"))
	req.Header.Set(vodka.Authorization, auth)
	he := ba(c).(*vodka.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, he.Code())
	assert.Equal(t, Basic+" realm=Restricted", rec.Header().Get(vodka.WWWAuthenticate))

	// Empty Authorization header
	req.Header.Set(vodka.Authorization, "")
	he = ba(c).(*vodka.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, he.Code())
	assert.Equal(t, Basic+" realm=Restricted", rec.Header().Get(vodka.WWWAuthenticate))

	// Invalid Authorization header
	auth = base64.StdEncoding.EncodeToString([]byte("invalid"))
	req.Header.Set(vodka.Authorization, auth)
	he = ba(c).(*vodka.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, he.Code())
	assert.Equal(t, Basic+" realm=Restricted", rec.Header().Get(vodka.WWWAuthenticate))

	// WebSocket
	c.Request().Header.Set(vodka.Upgrade, vodka.WebSocket)
	assert.NoError(t, ba(c))
}
