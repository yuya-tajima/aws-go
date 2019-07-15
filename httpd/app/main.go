package main

import (
	_ "fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/yuya-tajima/aws-go/httpd/app/controllers"
	"github.com/yuya-tajima/aws-go/httpd/app/template"
	"github.com/yuya-tajima/aws-go/httpd/app/auth"
	_"github.com/yuya-tajima/aws-go/httpd/app/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(auth.GenerateRandom(32)))))

	e.Renderer = template.GetTemplate()

	e.Static("/", "assets")

	//e.Use(_middleware.Author)

	e.GET("/", controllers.Index)
	e.GET("/ec2/desc", controllers.DescEc2)
	e.POST("/auth", controllers.Auth)

	e.POST("/ec2/start", controllers.StartEc2)
	e.POST("/ec2/stop", controllers.StopEc2)

	e.Logger.Fatal(e.Start(":8080"))
}
