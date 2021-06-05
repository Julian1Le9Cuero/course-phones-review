package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type MyServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Return a pointer to MyServer instance
	return &MyServer{s}
}

// Method to run server
func (s *MyServer) Run() {
	log.Fatal(s.server.ListenAndServe())
}
