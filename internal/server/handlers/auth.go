package handlers

import (
	"go-templ-htmx/web/views/auth"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func HandleLogin(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["foo"] = "bar"
	sess.Values["user"] = c.QueryParam("user")
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return Render(auth.Login(), c)
}
