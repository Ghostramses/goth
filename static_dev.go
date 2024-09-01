//go:build dev
// +build dev

package main

import "github.com/labstack/echo/v4"

func servePublic(app *echo.Echo) {
	app.Static("/public", "frontend/public")
}
