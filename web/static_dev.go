//go:build dev
// +build dev

package web

import "github.com/labstack/echo/v4"

func ServePublic(app *echo.Echo) {
	app.Static("/public", "web/assets")
}
