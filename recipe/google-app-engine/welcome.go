package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/insionng/vodka"
)

type (
	Template struct {
		templates *template.Template
	}
)

func init() {
	t := &Template{
		templates: template.Must(template.ParseFiles("templates/welcome.html")),
	}
	e.SetRenderer(t)
	e.GET("/welcome", welcome)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c vodka.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func welcome(c vodka.Context) error {
	return c.Render(http.StatusOK, "welcome", "Joe")
}
