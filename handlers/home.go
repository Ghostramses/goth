package handlers

import (
	"go-templ-htmx/frontend/views/home"

	"github.com/labstack/echo/v4"
)

func HandleHome(c echo.Context) error {
	return Render(home.Index(), c)
}
