## Subdomains

`server.go`

```go
package main

import (
	"net/http"

	"github.com/insionng/vodka"
	mw "github.com/insionng/vodka/middleware"
)

type Hosts map[string]http.Handler

func (h Hosts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := h[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	// Host map
	hosts := make(Hosts)

	//-----
	// API
	//-----

	api := vodka.New()
	api.Use(mw.Logger())
	api.Use(mw.Recover())

	hosts["api.localhost:1323"] = api

	api.Get("/", func(c *vodka.Context) error {
		return c.String(http.StatusOK, "API")
	})

	//------
	// Blog
	//------

	blog := vodka.New()
	blog.Use(mw.Logger())
	blog.Use(mw.Recover())

	hosts["blog.localhost:1323"] = blog

	blog.Get("/", func(c *vodka.Context) error {
		return c.String(http.StatusOK, "Blog")
	})

	//---------
	// Website
	//---------

	site := vodka.New()
	site.Use(mw.Logger())
	site.Use(mw.Recover())

	hosts["localhost:1323"] = site

	site.Get("/", func(c *vodka.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	http.ListenAndServe(":1323", hosts)
}
```

## [Source Code](https://github.com/insionng/vodka/blob/master/recipes/subdomains)
