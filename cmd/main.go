package main

import (
	"context"
	"go-templ-htmx/internal/server"
	"go-templ-htmx/internal/storage/database"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()

	db, err := database.NewPostgres(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	server := server.NewServer(db)
	if err := server.ListenAndServe(); err != nil {
		log.Println("Cannot start server")
		log.Fatalln(err)
	}
}
