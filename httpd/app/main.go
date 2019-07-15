package main

import (
	_ "fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yuya-tajima/aws-go/httpd/app/auth"
	"github.com/yuya-tajima/aws-go/httpd/app/controllers"
	mw "github.com/yuya-tajima/aws-go/httpd/app/middleware"
	"github.com/yuya-tajima/aws-go/httpd/app/template"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(auth.GenerateRandom(32)))))

	e.Renderer = template.GetTemplate()

	e.Static("/", "assets")

	e.Use(mw.Author)

	e.GET("/", controllers.Index)
	e.GET("/ec2/desc", controllers.DescEc2)
	e.POST("/auth", controllers.Auth)

	e.POST("/ec2/start", controllers.StartEc2)
	e.POST("/ec2/stop", controllers.StopEc2)

	e.Logger.Fatal(e.Start(":8080"))
}
