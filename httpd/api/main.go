package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/ec2/:operation", Ec2PostOperations)
	e.GET("/ec2/:operation", Ec2GetOperations)

	e.Logger.Fatal(e.Start(":8081"))
}
