package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port    string
	Handler http.Handler
}

func (s *Server) Run() {
	log.Printf("Running s on Port %s\n", s.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Port),
		Handler: s.Handler,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
