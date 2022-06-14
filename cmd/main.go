package main

import (
	"log"

	"github.com/Neko2h/shortener/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.NewApp()
}
