package server

import (
	"batch-v1/internal/server/routes"
	"log"
	"net/http"
)

func Server() {
	// Start HTTP server with route handler
	r := routes.Router()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
