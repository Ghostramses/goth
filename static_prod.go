//go:build !dev
// +build !dev

package main

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed frontend/public/*
var publicFs embed.FS

func servePublic(app *echo.Echo) {
	app.StaticFS("/public", publicFs)
}
