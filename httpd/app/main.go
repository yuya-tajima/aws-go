package main

import (
	_ "fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type metaData struct {
	Title string
}

func main() {

	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		var meta metaData
		meta.Title = "Test"

		return c.Render(http.StatusOK, "index", meta)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
