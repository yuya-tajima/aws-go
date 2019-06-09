package main

import (
	_ "fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuya-tajima/aws-go/aws"
)

var _aws *aws.Aws

func main() {

	e := echo.New()

	_aws = &aws.Aws{}
	_aws.SetSession("tajima")
	_aws.SetEc2Client()

	e.GET("/ec2/desc", func(c echo.Context) error {
		result, _ := _aws.Ec2.GetInstances("")

		var sCode int

		sCode = http.StatusOK

		return c.JSON(sCode, result)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
