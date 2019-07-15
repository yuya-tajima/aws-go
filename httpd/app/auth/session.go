package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/yuya-tajima/aws-go/httpd/app/util"
)

const sessKey = "auth_user"
const authorized = "complete"

const (
	OK = iota
	NG
)

type Context struct {
	echo.Context
	state int
}

func NewContext(c echo.Context) *Context {
	return &Context{c, NG}
}

func (c *Context) getSession() (sess *sessions.Session, err error) {
	sess, err = session.Get(sessKey, c)
	return
}

func (c *Context) DeleteAuth() {
	sess, err := c.getSession()
	if err == nil {
		delete(sess.Values, authorized)
		sess.Save(c.Request(), c.Response())
	}
}

func (c *Context) AuthJSON() error {

	status := http.StatusUnauthorized
	body := util.AuthErrorJSON(fmt.Errorf("%s", "Authentication failed"))

	if c.state == OK {
		status = http.StatusOK
		body = util.SuccessJSON()
	}

	return c.JSONBlob(status, body)
}

func (c *Context) Auth() error {

	c.state = NG

	l := new(Login)
	if err := c.Bind(l); err != nil {
		return err
	}

	if ok := AuthUser(l.Id, l.Pwd); !ok {
		return fmt.Errorf("%s", "error")
	}

	c.state = OK

	return nil
}

func (c *Context) SetAuthSession() {
	sess, _ := c.getSession()
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1,
		HttpOnly: true,
	}
	sess.Values[authorized] = true
	sess.Save(c.Request(), c.Response())
}

func (c *Context) IsAuthUser() bool {
	sess, err := c.getSession()
	if err != nil {
		return false
	}

	ok := sess.Values[authorized]

	if ok != nil && ok.(bool) {
		c.state = OK
		return true
	}

	return false
}
