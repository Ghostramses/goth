package server

import (
	"fmt"
	"go-templ-htmx/internal/server/handlers"
	imiddleware "go-templ-htmx/internal/server/middleware"
	"go-templ-htmx/internal/storage/database"
	"go-templ-htmx/web"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	db  database.DB
	app *echo.Echo
}

func NewServer(db database.DB) *Server {
	app := echo.New()

	// app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(imiddleware.AuthRequired)
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("testing"))))

	web.ServePublic(app)

	app.GET("/", handlers.HandleHome)
	app.GET("/login", handlers.HandleLogin)

	return &Server{db, app}
}

func (s *Server) ListenAndServe() error {

	listenAddr := os.Getenv("LISTEN_ADDR")
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		listenPort = "8080"
	}

	return s.app.Start(fmt.Sprintf("%s:%s", listenAddr, listenPort))
}
