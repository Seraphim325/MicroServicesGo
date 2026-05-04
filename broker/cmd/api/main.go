package main

import (
	"broker/cmd/api/router"
	"broker/cmd/api/server"
)

func main() {
	r := router.Router{}
	srv := server.Server{
		Port:    "8080",
		Handler: r.Handler(),
	}

	srv.Run()
}
