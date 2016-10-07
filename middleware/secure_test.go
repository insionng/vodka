package middleware

import (
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	"github.com/stretchr/testify/assert"
)

func TestSecure(t *testing.T) {
	e := vodka.New()
	req := test.NewRequest(vodka.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := func(c vodka.Context) error {
		return c.String(http.StatusOK, "test")
	}

	// Default
	Secure()(h)(c)
	assert.Equal(t, "1; mode=block", rec.Header().Get(vodka.HeaderXXSSProtection))
	assert.Equal(t, "nosniff", rec.Header().Get(vodka.HeaderXContentTypeOptions))
	assert.Equal(t, "SAMEORIGIN", rec.Header().Get(vodka.HeaderXFrameOptions))
	assert.Equal(t, "", rec.Header().Get(vodka.HeaderStrictTransportSecurity))
	assert.Equal(t, "", rec.Header().Get(vodka.HeaderContentSecurityPolicy))

	// Custom
	req.Header().Set(vodka.HeaderXForwardedProto, "https")
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	SecureWithConfig(SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	})(h)(c)
	assert.Equal(t, "", rec.Header().Get(vodka.HeaderXXSSProtection))
	assert.Equal(t, "", rec.Header().Get(vodka.HeaderXContentTypeOptions))
	assert.Equal(t, "", rec.Header().Get(vodka.HeaderXFrameOptions))
	assert.Equal(t, "max-age=3600; includeSubdomains", rec.Header().Get(vodka.HeaderStrictTransportSecurity))
	assert.Equal(t, "default-src 'self'", rec.Header().Get(vodka.HeaderContentSecurityPolicy))
}
