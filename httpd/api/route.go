package main

import (
	"github.com/labstack/echo"
)

func ec2Route(e *echo.Echo) {
	ec2 := e.Group("/ec2", ec2Client)
	ec2.POST("/:operation", ec2PostOperations)
	ec2.GET("/:operation", ec2GetOperations)
}
