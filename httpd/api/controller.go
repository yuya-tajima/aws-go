package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuya-tajima/aws-go/aws"
	"github.com/yuya-tajima/aws-go/aws/util"
)

type awsContext struct {
	echo.Context
	_aws *aws.Aws
}

func ec2PostOperations(c echo.Context) error {

	operation := c.Param("operation")

	cc := c.(*awsContext)

	var err error
	switch operation {
	case "start":
		err = cc.ec2Start()
	case "stop":
		err = cc.ec2Stop()
	case "reboot":
		err = cc.ec2Reboot()
	default:
		err = echo.NotFoundHandler(c)
	}

	return err
}

func ec2GetOperations(c echo.Context) error {

	operation := c.Param("operation")

	cc := c.(*awsContext)

	var err error
	switch operation {
	case "desc":
		err = cc.ec2Desc()
	default:
		err = echo.NotFoundHandler(c)
	}

	return err
}

func (a *awsContext) ec2Reboot() error {

	_aws := a._aws
	c := a.Context

	insId := c.FormValue("ins-id")

	result, err := _aws.Ec2.Reboot(insId, false)

	if err != nil {
		errCode, errRes := util.GetErrorResponse(err)
		return c.JSON(errCode, errRes)
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func (a *awsContext) ec2Start() error {

	_aws := a._aws
	c := a.Context

	insId := c.FormValue("ins-id")

	result, err := _aws.Ec2.Start(insId, false)

	if err != nil {
		errCode, errRes := util.GetErrorResponse(err)
		return c.JSON(errCode, errRes)
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func (a *awsContext) ec2Stop() error {

	_aws := a._aws
	c := a.Context

	insId := c.FormValue("ins-id")

	result, err := _aws.Ec2.Stop(insId, false)

	if err != nil {
		errCode, errRes := util.GetErrorResponse(err)
		return c.JSON(errCode, errRes)
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func (a *awsContext) ec2Desc() error {
	_aws := a._aws
	c := a.Context

	result, err := _aws.Ec2.GetInstances("")

	if err != nil {
		errCode, errRes := util.GetErrorResponse(err)
		return c.JSON(errCode, errRes)
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
