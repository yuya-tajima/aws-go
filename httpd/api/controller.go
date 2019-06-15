package main

import (
	_ "fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuya-tajima/aws-go/aws/util"
)

func Ec2PostOperations(c echo.Context) error {

	operation := c.Param("operation")

	var err error
	switch operation {
	case "start":
		err = ec2Start(c)
	case "stop":
		err = ec2Stop(c)
	default:
		err = echo.NotFoundHandler(c)
	}

	return err
}

func Ec2GetOperations(c echo.Context) error {

	operation := c.Param("operation")

	var err error
	switch operation {
	case "desc":
		err = ec2Desc(c)
	default:
		err = echo.NotFoundHandler(c)
	}

	return err
}

func ec2Start(c echo.Context) error {

	_aws := GetEc2Client(c)

	insId := c.FormValue("ins-id")

	result, err := _aws.Ec2.Start(insId, false)

	if err != nil {
		errCode, errRes := util.GetErrorResponse(err)
		return c.JSON(errCode, errRes)
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func ec2Stop(c echo.Context) error {

	_aws := GetEc2Client(c)

	insId := c.FormValue("ins-id")

	result, err := _aws.Ec2.Stop(insId, false)

	if err != nil {
		errCode, errRes := util.GetErrorResponse(err)
		return c.JSON(errCode, errRes)
	} else {
		return c.JSON(http.StatusOK, result)
	}

}

func ec2Desc(c echo.Context) error {

	_aws := GetEc2Client(c)

	result, err := _aws.Ec2.GetInstances("")

	if err != nil {
		errCode, errRes := util.GetErrorResponse(err)
		return c.JSON(errCode, errRes)
	} else {
		return c.JSON(http.StatusOK, result)
	}

}
