package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type userDetails struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"work"`
}

func profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(userDetails{
		Name:       "Emma!",
		Email:      "emma@gmail.com",
		Occupation: "Software Engineer",
	}); err != nil {
		log.Panicf("write to this: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allow", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok"}); err != nil {
		log.Printf("write health response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheck)
	mux.HandleFunc("/profile", profile)

	log.Println("server listening on port:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello from Kestrel")
}
