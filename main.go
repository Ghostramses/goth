package main

import (
	"fmt"
	"go-templ-htmx/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	app := echo.New()
	servePublic(app)
	app.GET("/", handlers.HandleHome)
	app.GET("/login", handlers.HandleLogin)

	listenAddr := os.Getenv("LISTEN_ADDR")
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		listenPort = "8080"
	}
	log.Fatalln(app.Start(fmt.Sprintf("%s:%s", listenAddr, listenPort)))
}
