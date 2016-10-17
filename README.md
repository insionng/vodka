# Vodka V2+

    由于Echo V3不再支持fasthttp, 于是我将以Vodka的名义自行维护Echo V2的后续开发，如果你也在使用我的这个版本欢迎留言交流.


#### Fast and unfancy HTTP server framework for Go (Golang). Up to 10x faster than the rest.

## Feature Overview

- Optimized HTTP router which smartly prioritize routes
- Build robust and scalable RESTful APIs
- Run with standard HTTP server or FastHTTP server
- Group APIs
- Extensible middleware framework
- Define middleware at root, group or route level
- Data binding for JSON, XML and form payload
- Handy functions to send variety of HTTP responses
- Centralized HTTP error handling
- Template rendering with any template engine
- Define your format for the logger
- Highly customizable

## Performance

- Environment:
	- Go 1.7.1
	- wrk 4.2.0
	- Memory 16 GB
    - Processor Intel® Xeon® CPU E3-1231 v3 @ 3.40GHz × 8

    Simple Test:
    ```go
    package main

    import (
	    "net/http"

        "github.com/insionng/vodka"
	    "github.com/insionng/vodka/engine/fasthttp"
    )

    func main() {
	    v := vodka.New()
        v.GET("/", HelloHandler)
	    v.Run(fasthttp.New(":1987"))
    }

    func HelloHandler(self vodka.Context) error {
	    return self.String(http.StatusOK, "Hello, World!")
    }

    ```


    ```sh
    wrk -t8 -c400 -d20s http://localhost:1987
    Running 20s test @ http://localhost:1987
    8 threads and 400 connections
    Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.06ms    1.15ms  38.67ms   92.60%
    Req/Sec    54.87k     6.99k   77.05k    75.69%
    8747188 requests in 20.05s, 1.21GB read
    Requests/sec: 436330.95
    Transfer/sec:     61.59MB
    ```

## 快速开始

### 安装

在安装之前确认你已经安装了Go语言. Go语言安装请访问 [install instructions](http://golang.org/doc/install.html).

Vodka is developed and tested using Go `1.7.x`+

```sh
$ go get -u github.com/insionng/vodka
```


### Hello, World!

Create `server.go`

```go
package main

import (
	"net/http"
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/fasthttp"
)

func main() {
	v := vodka.New()
	v.GET("/", func(self vodka.Context) error {
		return self.String(http.StatusOK, "Hello, World!")
	})
	v.Run(fasthttp.New(":1987"))
}
```

Start server

```sh
$ go run server.go
```

Browse to [http://localhost:1987](http://localhost:1987) and you should see
Hello, World! on the page.

### Routing

```go
v.POST("/users", saveUser)
v.GET("/users/:id", getUser)
v.PUT("/users/:id", updateUser)
v.DELETE("/users/:id", deleteUser)
```

### Path Parameters

```go
func getUser(self vodka.Context) error {
	// User ID from path `users/:id`
	id := self.Param("id")
}
```

### Query Parameters

`/show?team=x-men&member=wolverine`

```go
func show(c vodka.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
}
```

### Form `application/x-www-form-urlencoded`

`POST` `/save`

name | value
:--- | :---
name | Joe Smith
email | vodka@yougam.com


```go
func save(c vodka.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
}
```

### Form `multipart/form-data`

`POST` `/save`

name | value
:--- | :---
name | Joe Smith
email | joe@yougam.com
avatar | avatar

```go
func save(c vodka.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you!</b>")
}
```

### Handling Request

- Bind `JSON` or `XML` or `form` payload into Go struct based on `Content-Type` request header.
- Render response as `JSON` or `XML` with status code.

```go
type User struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
}

e.POST("/users", func(c vodka.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
	// or
	// return c.XML(http.StatusCreated, u)
})
```

### Static Content

Server any file from static directory for path `/static/*`.

```go
e.Static("/static", "static")
```

##### [Learn More](https://github.com/insionng/vodka/blob/master/docs/guide/static-files)

### [Template Rendering](https://github.com/insionng/vodka/blob/master/docs/guide/templates)

### Middleware

```go
// Root level middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// Group level middleware
g := e.Group("/root")
g.Use(middleware.BasicAuth(func(username, password string) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))

// Route level middleware
track := func(next vodka.HandlerFunc) vodka.HandlerFunc {
	return func(c vodka.Context) error {
		println("request to /users")
		return next(c)
	}
}
e.GET("/users", func(c vodka.Context) error {
	return c.String(http.StatusOK, "/users")
}, track)
```

#### Built-in Middleware

Middleware | Description
:--- | :---
[BodyLimit](https://github.com/insionng/vodka/blob/master/docs/middleware/body-limit) | Limit request body
[Logger](https://github.com/insionng/vodka/blob/master/docs/middleware/logger) | Log HTTP requests
[Recover](https://github.com/insionng/vodka/blob/master/docs/middleware/recover) | Recover from panics
[Gzip](https://github.com/insionng/vodka/blob/master/docs/middleware/gzip) | Send gzip HTTP response
[BasicAuth](https://github.com/insionng/vodka/blob/master/docs/middleware/basic-auth) | HTTP basic authentication
[JWTAuth](https://github.com/insionng/vodka/blob/master/docs/middleware/jwt) | JWT authentication
[Secure](https://github.com/insionng/vodka/blob/master/docs/middleware/secure) | Protection against attacks
[CORS](https://github.com/insionng/vodka/blob/master/docs/middleware/cors) | Cross-Origin Resource Sharing
[CSRF](https://github.com/insionng/vodka/blob/master/docs/middleware/csrf) | Cross-Site Request Forgery
[Static](https://github.com/insionng/vodka/blob/master/docs/middleware/static) | Serve static files
[HTTPSRedirect](https://github.com/insionng/vodka/blob/master/docs/middleware/redirect#httpsredirect-middleware) | Redirect HTTP requests to HTTPS
[HTTPSWWWRedirect](https://github.com/insionng/vodka/blob/master/docs/middleware/redirect#httpswwwredirect-middleware) | Redirect HTTP requests to WWW HTTPS
[WWWRedirect](https://github.com/insionng/vodka/blob/master/docs/middleware/redirect#wwwredirect-middleware) | Redirect non WWW requests to WWW
[NonWWWRedirect](https://github.com/insionng/vodka/blob/master/docs/middleware/redirect#nonwwwredirect-middleware) | Redirect WWW requests to non WWW
[AddTrailingSlash](https://github.com/insionng/vodka/blob/master/docs/middleware/trailing-slash#addtrailingslash-middleware) | Add trailing slash to the request URI
[RemoveTrailingSlash](https://github.com/insionng/vodka/blob/master/docs/middleware/trailing-slash#removetrailingslash-middleware) | Remove trailing slash from the request URI
[MethodOverride](https://github.com/insionng/vodka/blob/master/docs/middleware/method-override) | Override request method

##### [Learn More](https://github.com/insionng/vodka/blob/master/docs/middleware/overview)

#### Third-party Middleware

Middleware | Description
:--- | :---
[vodkaperm](https://github.com/xyproto/vodkaperm) | Keeping track of users, login states and permissions.
[vodkapprof](https://github.com/vodka-contrib/vodkapprof) | Adapt net/http/pprof to vodka.


### Need help?

- [QQ Group] Vodka/Vodka Web 框架群号 242851426
- [Open an issue](https://github.com/insionng/vodka/issues/new)

## Support Us

- :star: the project
- :earth_americas: spread the word
- [Contribute](#contribute) to the project

## Contribute

**Use issues for everything**

- Report issues
- Discuss on chat before sending a pull request
- Suggest new features or enhancements
- Improve/fix documentation


## Vodka System

Community created packages for Vodka

- [hello world](https://github.com/vodka-contrib/helloworld)


## Vodka Case

- [ZenPress](https://github.com/insionng/zenpress) - Cms/Blog System(just start-up)


## Donation
    BTC:1JHtavsBqBNGxpR4eLNwjYL9Vjbyr3Tw6T

## License
    MIT License


## QQ Group

    Vodka/Vodka Web 框架群号 242851426

    Golang编程(Go/Web/Nosql)群号 245386165

    Go语言编程(Golang/Web/Orm)群号 231956113

    Xorm & Golang群号 280360085

    Golang & Tango & Web群号 369240307

    Martini&Macaron 交流群 371440803
