package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yuya-tajima/aws-go/aws"
)

type awsContext struct {
	echo.Context
	_aws *aws.Aws
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ec2 := e.Group("/ec2", ec2Client)
	ec2.POST("/:operation", ec2PostOperations)
	ec2.GET("/:operation", ec2GetOperations)

	e.Logger.Fatal(e.Start(":8081"))
}
