package main

import (
	"github.com/labstack/echo"
	"github.com/yuya-tajima/aws-go/aws"
)

func GetEc2Client(c echo.Context) *aws.Aws {

	h := c.Request().Header

	_aws := &aws.Aws{}
	_aws.SetStaticSession(h.Get("Secret_Id"), h.Get("Secret_key"), "", "")

	cfg := &aws.Ec2Config{
		Region: h.Get("Region"),
	}

	_aws.SetEc2Client(cfg)

	return _aws
}
