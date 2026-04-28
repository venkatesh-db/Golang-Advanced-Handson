

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response model
type HealthResponse struct {
	Status string `json:"status"`
}

// Middleware: simple authorization check
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "Bearer production-token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()

	// Route with middleware
	mux.Handle("/health", authMiddleware(http.HandlerFunc(healthHandler)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server started on http://localhost:8080")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}