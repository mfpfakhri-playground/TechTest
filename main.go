package main

import (
	"log"
	"net/http"
	"os"
	"tech-test/database"
	"tech-test/routes"
	"time"
)

func main() {
	database.Initiation()
	serveHTTP()
}

func serveHTTP() {
	r := routes.Initiation()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	addr := "0.0.0.0:8080"
	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("could not serve http server on port 8080: %s", err.Error())
	}
}
