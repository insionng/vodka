package fasthttp

import (
	"bufio"
	"bytes"
	"net"
	"net/url"
	"testing"

	"github.com/insionng/vodka/engine/test"
	"github.com/insionng/vodka/libraries/gommon/log"
	fast "github.com/valyala/fasthttp"
)

type fakeAddr struct {
	addr string
	net.Addr
}

func (a fakeAddr) String() string {
	return a.addr
}

func TestRequest(t *testing.T) {
	ctx := new(fast.RequestCtx)
	url, _ := url.Parse("http://github.com/insionng/vodka")
	ctx.Init(&fast.Request{}, fakeAddr{addr: "127.0.0.1"}, nil)
	ctx.Request.Read(bufio.NewReader(bytes.NewBufferString(test.MultipartRequest)))
	ctx.Request.SetRequestURI(url.String())
	test.RequestTest(t, NewRequest(ctx, log.New("vodka")))
}
