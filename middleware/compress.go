package middleware

import (
	"bufio"
	"compress/gzip"
	"github.com/insionng/vodka"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
)

type (
	gzipWriter struct {
		io.Writer
		http.ResponseWriter
	}
)

func (w gzipWriter) Write(b []byte) (int, error) {
	if w.Header().Get(vodka.ContentType) == "" {
		w.Header().Set(vodka.ContentType, http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

func (w gzipWriter) Flush() error {
	return w.Writer.(*gzip.Writer).Flush()
}

func (w gzipWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (w *gzipWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

var writerPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(ioutil.Discard)
	},
}

// Gzip returns a middleware which compresses HTTP response using gzip compression
// scheme.
func Gzip() vodka.MiddlewareFunc {
	scheme := "gzip"

	return func(h vodka.HandlerFunc) vodka.HandlerFunc {
		return func(c *vodka.Context) error {
			c.Response().Header().Add(vodka.Vary, vodka.AcceptEncoding)
			if strings.Contains(c.Request().Header.Get(vodka.AcceptEncoding), scheme) {
				w := writerPool.Get().(*gzip.Writer)
				w.Reset(c.Response().Writer())
				defer func() {
					w.Close()
					writerPool.Put(w)
				}()
				gw := gzipWriter{Writer: w, ResponseWriter: c.Response().Writer()}
				c.Response().Header().Set(vodka.ContentEncoding, scheme)
				c.Response().SetWriter(gw)
			}
			if err := h(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
