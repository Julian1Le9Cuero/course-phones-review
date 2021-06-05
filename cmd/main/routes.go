package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()
	// Middlewares globales (que se usarán en todas las rutas)
	mux.Use(
		middleware.Logger,    // Log every http request
		middleware.Recoverer, // Recover if a panic occurs
	)

	// Routes
	// Add new smartphone
	mux.Post("/smartphones", nil)
	// Get message
	mux.Get("/hello", helloHandler)

	return mux
}

// Handler function for routes
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set writes new header if not exists, otherwise it overwrites it
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "julian")

	res := map[string]interface{}{"message": "Hello world"}
	// _ = Significa que si se arroja un error se ignorará
	_ = json.NewEncoder(w).Encode(res)

}
