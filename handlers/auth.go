package handlers

import (
	"go-templ-htmx/frontend/views/auth"

	"github.com/labstack/echo/v4"
)

func HandleLogin(c echo.Context) error {
	return Render(auth.Login(), c)
}
