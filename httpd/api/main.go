package main

import (
	"bytes"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuya-tajima/aws-go/aws"
)

var _aws *aws.Aws

func main() {

	e := echo.New()
	var buffer bytes.Buffer

	_aws = &aws.Aws{}
	_aws.SetSession("tajima")
	_aws.SetEc2Client()

	e.GET("/ec2/desc", func(c echo.Context) error {
		buffer.Reset()
		result, _ := _aws.GetInstances("")
		if len(result) > 0 {
			for _, i := range result {
				buffer.WriteString(*i.InstanceId)
			}
		}

		return c.String(http.StatusOK, buffer.String())

	})

	e.Logger.Fatal(e.Start(":8080"))

}
