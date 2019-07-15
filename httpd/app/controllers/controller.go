package controllers

import (
	_"fmt"
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/yuya-tajima/aws-go/httpd/app/client"
	"github.com/yuya-tajima/aws-go/httpd/app/auth"
	"github.com/yuya-tajima/aws-go/httpd/api/data"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func DescEc2(c echo.Context) error {
	url := "http://localhost:8081/ec2/desc"
	body, status := client.GetRequest(url)
	return c.JSONBlob(status, body)
}

func StartEc2(c echo.Context) error {
	i := new(data.InsTance)

	if err := c.Bind(i); err != nil {
		return err
	}

	json := `{"ins_id":"` + i.InsID + `"}`

	url := "http://localhost:8081/ec2/start"

	body, status := client.PostRequest(url, []byte(json))
	return c.JSONBlob(status, body)
}

func StopEc2(c echo.Context) error {
	i := new(InsTance)
	if err := c.Bind(i); err != nil {
		return err
	}

	json := `{"ins_id":"` + i.InsID + `"}`

	url := "http://localhost:8081/ec2/stop"

	body, status := client.PostRequest(url, []byte(json))
	return c.JSONBlob(status, body)
}

func Auth(c echo.Context) error {

	cc := auth.NewContext(c)

	if ok := cc.IsAuthUser(); ! ok {
		if err := cc.Auth(); err != nil {
			return cc.AuthJSON()
		}
	}

	cc.SetAuthSession()

	return cc.AuthJSON()
}
