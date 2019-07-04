package main
import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	_ "github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/labstack/echo-contrib/session"
)

const profileName = "aws_api" 

var ( access_key string
	secret_key string
)

type cred struct {
	Access_key string `json:"access_key"`
	Secret_key string `json:"secret_key"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	err := setCredByEnv()
	if err != nil {
		fmt.Printf("NOTICE:%s\n", err)
		err = setCredByFile()
		if err != nil {
			fmt.Printf("ERROR:%s\n", err)
			os.Exit(1)
		}
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.Static("/", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/cred", func(c echo.Context) error {
		res := &cred{
			Access_key:access_key,
			Secret_key:secret_key,
		}
		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
