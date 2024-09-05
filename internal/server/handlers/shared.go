package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(component templ.Component, c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/html")
	return component.Render(c.Request().Context(), c.Response().Writer)
}
