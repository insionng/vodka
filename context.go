package vodka

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/insionng/vodka/engine"
	"github.com/insionng/vodka/log"

	"bytes"

	"github.com/insionng/vodka/context"
)

type (
	// Context represents the context of the current HTTP request. It holds request and
	// response objects, path, path parameters, data and registered handler.
	Context interface {
		context.Context

		// StdContext returns `context.Context`.
		StdContext() context.Context

		// SetStdContext sets `context.Context`.
		SetStdContext(context.Context)

		// Request returns `engine.Request` interface.
		Request() engine.Request

		// Request returns `engine.Response` interface.
		Response() engine.Response

		// Path returns the registered path for the handler.
		Path() string

		// SetPath sets the registered path for the handler.
		SetPath(string)

		// P returns path parameter by index.
		P(int) string

		// Param returns path parameter by name.
		Param(string) string

		// ParamNames returns path parameter names.
		ParamNames() []string

		// SetParamNames sets path parameter names.
		SetParamNames(...string)

		// ParamValues returns path parameter values.
		ParamValues() []string

		// SetParamValues sets path parameter values.
		SetParamValues(...string)

		// QueryParam returns the query param for the provided name. It is an alias
		// for `engine.URL#QueryParam()`.
		QueryParam(string) string

		// QueryParams returns the query parameters as map.
		// It is an alias for `engine.URL#QueryParams()`.
		QueryParams() map[string][]string

		// FormValue returns the form field value for the provided name. It is an
		// alias for `engine.Request#FormValue()`.
		FormValue(string) string

		// FormParams returns the form parameters as map.
		// It is an alias for `engine.Request#FormParams()`.
		FormParams() map[string][]string

		// FormFile returns the multipart form file for the provided name. It is an
		// alias for `engine.Request#FormFile()`.
		FormFile(string) (*multipart.FileHeader, error)

		// MultipartForm returns the multipart form.
		// It is an alias for `engine.Request#MultipartForm()`.
		MultipartForm() (*multipart.Form, error)

		// Cookie returns the named cookie provided in the request.
		// It is an alias for `engine.Request#Cookie()`.
		Cookie(string) (engine.Cookie, error)

		// SetCookie adds a `Set-Cookie` header in HTTP response.
		// It is an alias for `engine.Response#SetCookie()`.
		SetCookie(engine.Cookie)

		// Cookies returns the HTTP cookies sent with the request.
		// It is an alias for `engine.Request#Cookies()`.
		Cookies() []engine.Cookie

		// Get retrieves data from the context.
		Get(string) interface{}

		// Set saves data in the context.
		Set(string, interface{})

		// Bind binds the request body into provided type `i`. The default binder
		// does it based on Content-Type header.
		Bind(interface{}) error

		// Render renders a template with data and sends a text/html response with status
		// code. Templates can be registered using `Vodka.SetRenderer()`.
		Render(int, string, interface{}) error

		// HTML sends an HTTP response with status code.
		HTML(int, string) error

		// String sends a string response with status code.
		String(int, string) error

		// JSON sends a JSON response with status code.
		JSON(int, interface{}) error

		// JSONBlob sends a JSON blob response with status code.
		JSONBlob(int, []byte) error

		// JSONP sends a JSONP response with status code. It uses `callback` to construct
		// the JSONP payload.
		JSONP(int, string, interface{}) error

		// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
		// to construct the JSONP payload.
		JSONPBlob(int, string, []byte) error

		// XML sends an XML response with status code.
		XML(int, interface{}) error

		// XMLBlob sends a XML blob response with status code.
		XMLBlob(int, []byte) error

		// Blob sends a blob response with status code and content type.
		Blob(int, string, []byte) error

		// Stream sends a streaming response with status code and content type.
		Stream(int, string, io.Reader) error

		// File sends a response with the content of the file.
		File(string) error

		// Attachment sends a response from `io.ReaderSeeker` as attachment, prompting
		// client to save the file.
		Attachment(io.ReadSeeker, string) error

		// Inline sends a response from `io.ReaderSeeker` as inline, opening
		// the file in the browser.
		Inline(io.ReadSeeker, string) error

		// NoContent sends a response with no body and a status code.
		NoContent(int) error

		// Redirect redirects the request with status code.
		Redirect(int, string) error

		// Error invokes the registered HTTP error handler. Generally used by middleware.
		Error(err error)

		// Handler returns the matched handler by router.
		Handler() HandlerFunc

		// SetHandler sets the matched handler by router.
		SetHandler(HandlerFunc)

		// Logger returns the `Logger` instance.
		Logger() log.Logger

		// Vodka returns the `Vodka` instance.
		Vodka() *Vodka

		// ServeContent sends static content from `io.Reader` and handles caching
		// via `If-Modified-Since` request header. It automatically sets `Content-Type`
		// and `Last-Modified` response headers.
		ServeContent(io.ReadSeeker, string, time.Time) error

		// Reset resets the context after request completes. It must be called along
		// with `Vodka#AcquireContext()` and `Vodka#ReleaseContext()`.
		// See `Vodka#ServeHTTP()`
		Reset(engine.Request, engine.Response)
	}

	vodkaContext struct {
		context  context.Context
		request  engine.Request
		response engine.Response
		path     string
		pnames   []string
		pvalues  []string
		handler  HandlerFunc
		vodka     *Vodka
	}
)

const (
	indexPage = "index.html"
)

func (c *vodkaContext) StdContext() context.Context {
	return c.context
}

func (c *vodkaContext) SetStdContext(ctx context.Context) {
	c.context = ctx
}

func (c *vodkaContext) Deadline() (deadline time.Time, ok bool) {
	return c.context.Deadline()
}

func (c *vodkaContext) Done() <-chan struct{} {
	return c.context.Done()
}

func (c *vodkaContext) Err() error {
	return c.context.Err()
}

func (c *vodkaContext) Value(key interface{}) interface{} {
	return c.context.Value(key)
}

func (c *vodkaContext) Request() engine.Request {
	return c.request
}

func (c *vodkaContext) Response() engine.Response {
	return c.response
}

func (c *vodkaContext) Path() string {
	return c.path
}

func (c *vodkaContext) SetPath(p string) {
	c.path = p
}

func (c *vodkaContext) P(i int) (value string) {
	l := len(c.pnames)
	if i < l {
		value = c.pvalues[i]
	}
	return
}

func (c *vodkaContext) Param(name string) (value string) {
	l := len(c.pnames)
	for i, n := range c.pnames {
		if n == name && i < l {
			value = c.pvalues[i]
			break
		}
	}
	return
}

func (c *vodkaContext) ParamNames() []string {
	return c.pnames
}

func (c *vodkaContext) SetParamNames(names ...string) {
	c.pnames = names
}

func (c *vodkaContext) ParamValues() []string {
	return c.pvalues
}

func (c *vodkaContext) SetParamValues(values ...string) {
	c.pvalues = values
}

func (c *vodkaContext) QueryParam(name string) string {
	return c.request.URL().QueryParam(name)
}

func (c *vodkaContext) QueryParams() map[string][]string {
	return c.request.URL().QueryParams()
}

func (c *vodkaContext) FormValue(name string) string {
	return c.request.FormValue(name)
}

func (c *vodkaContext) FormParams() map[string][]string {
	return c.request.FormParams()
}

func (c *vodkaContext) FormFile(name string) (*multipart.FileHeader, error) {
	return c.request.FormFile(name)
}

func (c *vodkaContext) MultipartForm() (*multipart.Form, error) {
	return c.request.MultipartForm()
}

func (c *vodkaContext) Cookie(name string) (engine.Cookie, error) {
	return c.request.Cookie(name)
}

func (c *vodkaContext) SetCookie(cookie engine.Cookie) {
	c.response.SetCookie(cookie)
}

func (c *vodkaContext) Cookies() []engine.Cookie {
	return c.request.Cookies()
}

func (c *vodkaContext) Set(key string, val interface{}) {
	c.context = context.WithValue(c.context, key, val)
}

func (c *vodkaContext) Get(key string) interface{} {
	return c.context.Value(key)
}

func (c *vodkaContext) Bind(i interface{}) error {
	return c.vodka.binder.Bind(i, c)
}

func (c *vodkaContext) Render(code int, name string, data interface{}) (err error) {
	if c.vodka.renderer == nil {
		return ErrRendererNotRegistered
	}
	buf := new(bytes.Buffer)
	if err = c.vodka.renderer.Render(buf, name, data, c); err != nil {
		return
	}
	c.response.Header().Set(HeaderContentType, MIMETextHTMLCharsetUTF8)
	c.response.WriteHeader(code)
	_, err = c.response.Write(buf.Bytes())
	return
}

func (c *vodkaContext) HTML(code int, html string) (err error) {
	c.response.Header().Set(HeaderContentType, MIMETextHTMLCharsetUTF8)
	c.response.WriteHeader(code)
	_, err = c.response.Write([]byte(html))
	return
}

func (c *vodkaContext) String(code int, s string) (err error) {
	c.response.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	c.response.WriteHeader(code)
	_, err = c.response.Write([]byte(s))
	return
}

func (c *vodkaContext) JSON(code int, i interface{}) (err error) {
	b, err := json.Marshal(i)
	if c.vodka.Debug() {
		b, err = json.MarshalIndent(i, "", "  ")
	}
	if err != nil {
		return err
	}
	return c.JSONBlob(code, b)
}

func (c *vodkaContext) JSONBlob(code int, b []byte) (err error) {
	return c.Blob(code, MIMEApplicationJSONCharsetUTF8, b)
}

func (c *vodkaContext) JSONP(code int, callback string, i interface{}) (err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return c.JSONPBlob(code, callback, b)
}

func (c *vodkaContext) JSONPBlob(code int, callback string, b []byte) (err error) {
	c.response.Header().Set(HeaderContentType, MIMEApplicationJavaScriptCharsetUTF8)
	c.response.WriteHeader(code)
	if _, err = c.response.Write([]byte(callback + "(")); err != nil {
		return
	}
	if _, err = c.response.Write(b); err != nil {
		return
	}
	_, err = c.response.Write([]byte(");"))
	return
}

func (c *vodkaContext) XML(code int, i interface{}) (err error) {
	b, err := xml.Marshal(i)
	if c.vodka.Debug() {
		b, err = xml.MarshalIndent(i, "", "  ")
	}
	if err != nil {
		return err
	}
	return c.XMLBlob(code, b)
}

func (c *vodkaContext) XMLBlob(code int, b []byte) (err error) {
	if _, err = c.response.Write([]byte(xml.Header)); err != nil {
		return
	}
	return c.Blob(code, MIMEApplicationXMLCharsetUTF8, b)
}

func (c *vodkaContext) Blob(code int, contentType string, b []byte) (err error) {
	c.response.Header().Set(HeaderContentType, contentType)
	c.response.WriteHeader(code)
	_, err = c.response.Write(b)
	return
}

func (c *vodkaContext) Stream(code int, contentType string, r io.Reader) (err error) {
	c.response.Header().Set(HeaderContentType, contentType)
	c.response.WriteHeader(code)
	_, err = io.Copy(c.response, r)
	return
}

func (c *vodkaContext) File(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return ErrNotFound
	}
	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		file = filepath.Join(file, "index.html")
		f, err = os.Open(file)
		if err != nil {
			return ErrNotFound
		}
		if fi, err = f.Stat(); err != nil {
			return err
		}
	}
	return c.ServeContent(f, fi.Name(), fi.ModTime())
}

func (c *vodkaContext) Attachment(r io.ReadSeeker, name string) (err error) {
	return c.contentDisposition(r, name, "attachment")
}

func (c *vodkaContext) Inline(r io.ReadSeeker, name string) (err error) {
	return c.contentDisposition(r, name, "inline")
}

func (c *vodkaContext) contentDisposition(r io.ReadSeeker, name, dispositionType string) (err error) {
	c.response.Header().Set(HeaderContentType, ContentTypeByExtension(name))
	c.response.Header().Set(HeaderContentDisposition, fmt.Sprintf("%s; filename=%s", dispositionType, name))
	c.response.WriteHeader(http.StatusOK)
	_, err = io.Copy(c.response, r)
	return
}

func (c *vodkaContext) NoContent(code int) error {
	c.response.WriteHeader(code)
	return nil
}

func (c *vodkaContext) Redirect(code int, url string) error {
	if code < http.StatusMultipleChoices || code > http.StatusTemporaryRedirect {
		return ErrInvalidRedirectCode
	}
	c.response.Header().Set(HeaderLocation, url)
	c.response.WriteHeader(code)
	return nil
}

func (c *vodkaContext) Error(err error) {
	c.vodka.httpErrorHandler(err, c)
}

func (c *vodkaContext) Vodka() *Vodka {
	return c.vodka
}

func (c *vodkaContext) Handler() HandlerFunc {
	return c.handler
}

func (c *vodkaContext) SetHandler(h HandlerFunc) {
	c.handler = h
}

func (c *vodkaContext) Logger() log.Logger {
	return c.vodka.logger
}

func (c *vodkaContext) ServeContent(content io.ReadSeeker, name string, modtime time.Time) error {
	req := c.Request()
	res := c.Response()

	if t, err := time.Parse(http.TimeFormat, req.Header().Get(HeaderIfModifiedSince)); err == nil && modtime.Before(t.Add(1*time.Second)) {
		res.Header().Del(HeaderContentType)
		res.Header().Del(HeaderContentLength)
		return c.NoContent(http.StatusNotModified)
	}

	res.Header().Set(HeaderContentType, ContentTypeByExtension(name))
	res.Header().Set(HeaderLastModified, modtime.UTC().Format(http.TimeFormat))
	res.WriteHeader(http.StatusOK)
	_, err := io.Copy(res, content)
	return err
}

// ContentTypeByExtension returns the MIME type associated with the file based on
// its extension. It returns `application/octet-stream` incase MIME type is not
// found.
func ContentTypeByExtension(name string) (t string) {
	if t = mime.TypeByExtension(filepath.Ext(name)); t == "" {
		t = MIMEOctetStream
	}
	return
}

func (c *vodkaContext) Reset(req engine.Request, res engine.Response) {
	c.context = context.Background()
	c.request = req
	c.response = res
	c.handler = NotFoundHandler
}
