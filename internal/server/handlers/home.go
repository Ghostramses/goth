package handlers

import (
	"fmt"
	"go-templ-htmx/web/views/home"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func HandleHome(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	fmt.Println(sess.Values)

	return Render(home.Index(), c)
}
