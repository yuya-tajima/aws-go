package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/yuya-tajima/aws-go/httpd/app/util"
)

const sessKey    = "auth_user"
const authorized = "complete"


type Context struct {
	echo.Context
	state int
}

func NewContext (c echo.Context) *Context {
	return &Context{ c, 0 }
}

func (c *Context) AuthJSON () error {

	status := http.StatusUnauthorized
	body := util.AuthErrorJSON(fmt.Errorf("%s", "Authentication failed"))

	if (c.state == 1) {
		status = http.StatusOK
		body = util.ConvertBytes("{\"auth\":1}")
	}

	return c.JSONBlob(status, body)
}

func (c *Context) Auth () error {

	l := new(Login)
	if err := c.Bind(l); err != nil {
		return err
	}

	if ok := AuthUser(l.Id, l.Pwd); ! ok {
		return fmt.Errorf("%s", "error")
	}

	return nil
}

func (c *Context) SetAuthSession () {
	sess, _ := session.Get(sessKey, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1,
		HttpOnly: true,
	}
	sess.Values[authorized] = true
	sess.Save(c.Request(), c.Response())
}

func (c *Context) IsAuthUser () bool {
	sess, err := session.Get(sessKey, c);
	if err != nil {
		return false
	}

	ok := sess.Values[authorized]
	
	if ok != nil && ok.(bool) {
		c.state = 1
		return true
	}

	return false
}
