# Vodka

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
	- wrk 4.0.0
	- 16 GB, 8 Core

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
email | joe@example.com


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
email | joe@example.com
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

##### [Learn More](https://vodka.labstack.com/guide/static-files)

### [Template Rendering](https://vodka.labstack.com/guide/templates)

### Middleware

```go
// Root level middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// Group level middleware
g := e.Group("/admin")
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
[BodyLimit](https://vodka.labstack.com/middleware/body-limit) | Limit request body
[Logger](https://vodka.labstack.com/middleware/logger) | Log HTTP requests
[Recover](https://vodka.labstack.com/middleware/recover) | Recover from panics
[Gzip](https://vodka.labstack.com/middleware/gzip) | Send gzip HTTP response
[BasicAuth](https://vodka.labstack.com/middleware/basic-auth) | HTTP basic authentication
[JWTAuth](https://vodka.labstack.com/middleware/jwt) | JWT authentication
[Secure](https://vodka.labstack.com/middleware/secure) | Protection against attacks
[CORS](https://vodka.labstack.com/middleware/cors) | Cross-Origin Resource Sharing
[CSRF](https://vodka.labstack.com/middleware/csrf) | Cross-Site Request Forgery
[Static](https://vodka.labstack.com/middleware/static) | Serve static files
[HTTPSRedirect](https://vodka.labstack.com/middleware/redirect#httpsredirect-middleware) | Redirect HTTP requests to HTTPS
[HTTPSWWWRedirect](https://vodka.labstack.com/middleware/redirect#httpswwwredirect-middleware) | Redirect HTTP requests to WWW HTTPS
[WWWRedirect](https://vodka.labstack.com/middleware/redirect#wwwredirect-middleware) | Redirect non WWW requests to WWW
[NonWWWRedirect](https://vodka.labstack.com/middleware/redirect#nonwwwredirect-middleware) | Redirect WWW requests to non WWW
[AddTrailingSlash](https://vodka.labstack.com/middleware/trailing-slash#addtrailingslash-middleware) | Add trailing slash to the request URI
[RemoveTrailingSlash](https://vodka.labstack.com/middleware/trailing-slash#removetrailingslash-middleware) | Remove trailing slash from the request URI
[MethodOverride](https://vodka.labstack.com/middleware/method-override) | Override request method

##### [Learn More](https://vodka.labstack.com/middleware/overview)

#### Third-party Middleware

Middleware | Description
:--- | :---
[vodkaperm](https://github.com/xyproto/vodkaperm) | Keeping track of users, login states and permissions.
[vodkapprof](https://github.com/mtojek/vodkapprof) | Adapt net/http/pprof to labstack/vodka.


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



## Credits
- [Vishal Rana](https://github.com/vishr) - Echo Author
- [Nitin Rana](https://github.com/nr17) - Consultant
- [Contributors](https://github.com/labstack/echo/graphs/contributors)
- [Insion Ng](https://github.com/insionng) - Vodka Author


## License
    MIT License


## QQ Group

    Vodka/Vodka Web 框架群号 242851426

    Golang编程(Go/Web/Nosql)群号 245386165

    Go语言编程(Golang/Web/Orm)群号 231956113

    Xorm & Golang群号 280360085

    Golang & Tango & Web群号 369240307

    Martini&Macaron 交流群 371440803
