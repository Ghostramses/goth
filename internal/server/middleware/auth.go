package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().RequestURI == "/login" {
			return next(c)
		}

		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		if sess.IsNew {
			return c.Redirect(http.StatusFound, "/login")

		}
		return next(c)
	}
}
