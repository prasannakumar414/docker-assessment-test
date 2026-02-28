package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type infoResponse struct {
	Email string `json:"email"`
}

type healthResponse struct {
	Status string `json:"status"`
}

func main() {
	email := os.Getenv("EMAIL")
	if email == "" {
		email = "test@example.com"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/info", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, infoResponse{Email: email})
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
	})

	addr := ":8080"
	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
