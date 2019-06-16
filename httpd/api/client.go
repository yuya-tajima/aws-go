package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/yuya-tajima/aws-go/aws"
)

func getEc2Client(c echo.Context) (*aws.Aws, error) {

	var err error

	h := c.Request().Header

	region := h.Get("Region")

	if region == "" {
		err = fmt.Errorf("Region does not exist")
	}

	id := h.Get("Secret_Id")

	if id == "" {
		err = fmt.Errorf("Secret_Id does not exist")
	}

	key := h.Get("Secret_key")

	if key == "" {
		err = fmt.Errorf("Secret_key does not exist")
	}

	_aws := &aws.Aws{}

	if err == nil {
		_aws.SetStaticSession(id, key, "", "")

		cfg := &aws.Ec2Config{
			Region: region,
		}

		_aws.SetEc2Client(cfg)
	}

	return _aws, err
}
