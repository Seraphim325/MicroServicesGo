package main

import (
	"broker/cmd/api/router"
	"broker/cmd/api/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadConfig()

	r := router.Router{}
	srv := server.Server{
		Port:    os.Getenv("BROKER_PORT"),
		Handler: r.Handler(),
	}

	srv.Run()
}

func loadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}
}
