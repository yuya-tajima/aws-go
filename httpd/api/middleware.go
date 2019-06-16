package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func ec2Client(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_aws, err := getEc2Client(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		cc := &awsContext{c, _aws}

		return next(cc)
	}
}
