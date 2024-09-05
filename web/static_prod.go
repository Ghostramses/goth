//go:build !dev
// +build !dev

package web

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed assets/*
var publicFs embed.FS

func ServePublic(app *echo.Echo) {
	app.StaticFS("/public", publicFs)
}
