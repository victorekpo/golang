package routes

import (
	"batch-v1/internal/server/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	// r.Use(middleware.AuthMiddleware)

	// Routes
	//r.HandleFunc("/api/health", handlers.HealthHandler).Methods("GET")
	//r.HandleFunc("/api/resource", handlers.CreateResourceHandler).Methods("POST")
	//r.HandleFunc("/api/resource/{id}", handlers.UpdateResourceHandler).Methods("PATCH")
	//r.HandleFunc("/ws", handlers.WebSocketHandler)

	return r
}
