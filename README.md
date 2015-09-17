# Vodka

Vodka是一个修改自Echo的强大Go语言web框架,仅作为我个人自用,主要为了解决依赖问题以及改造适应我自身使用，如果你也在使用我的改进版本欢迎留言交流.

## 安装

在安装之前确认你已经安装了Go语言. Go语言安装请访问 [install instructions](http://golang.org/doc/install.html). 

安装 vodka:

    go get github.com/insionng/vodka

A fast and unfancy micro web framework for Golang.


## Features

- Fast HTTP router which smartly prioritize routes.
- Extensible middleware, supports:
	- `vodka.MiddlewareFunc`
	- `func(vodka.HandlerFunc) vodka.HandlerFunc`
	- `vodka.HandlerFunc`
	- `func(*vodka.Context) error`
	- `func(http.Handler) http.Handler`
	- `http.Handler`
	- `http.HandlerFunc`
	- `func(http.ResponseWriter, *http.Request)`
- Extensible handler, supports:
    - `vodka.HandlerFunc`
    - `func(*vodka.Context) error`
    - `http.Handler`
    - `http.HandlerFunc`
    - `func(http.ResponseWriter, *http.Request)`
- Sub-router/Groups
- Handy functions to send variety of HTTP response:
    - HTML
    - HTML via templates
    - String 
    - JSON
    - JSONP
    - XML
    - File
    - Status
    - Redirect
    - Error
- Build-in support for:
	- Favicon
	- Index file
	- Static files
	- WebSocket
- Centralized HTTP error handling.
- Customizable HTTP request binding function.
- Customizable HTTP response rendering function, allowing you to use any HTML template engine.

## Vodka System

Community created packages for Vodka

- [hello world](https://github.com/vodka-contrib/helloworld)


## License
MIT License



