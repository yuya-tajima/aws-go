package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuya-tajima/aws-go/aws"
	"github.com/yuya-tajima/aws-go/aws/util"
)

var _aws *aws.Aws

func main() {

	e := echo.New()

	e.GET("/ec2/desc", func(c echo.Context) error {

		h := c.Request().Header

		_aws = &aws.Aws{}
		_aws.SetStaticSession(h.Get("Secret_Id"), h.Get("Secret_key"), "", "")

		cfg := &aws.Ec2Config{
			Region: "ap-northeast-1",
		}

		_aws.SetEc2Client(cfg)

		result, err := _aws.Ec2.GetInstances("")

		if err != nil {
			errCode, errRes := util.GetErrorResponse(err)
			return c.JSON(errCode, errRes)
		} else {
			return c.JSON(http.StatusOK, result)
		}
	})

	e.Logger.Fatal(e.Start(":8081"))
}
