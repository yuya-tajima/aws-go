package middleware

import (
	
	"github.com/labstack/echo/v4"
	"github.com/yuya-tajima/aws-go/httpd/app/auth"
)

func Author(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := auth.NewContext(c)
		if ok := cc.IsAuthUser(); ! ok {
			return cc.AuthJSON()
		}

		return next(cc)
	}
}
