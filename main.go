package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chan-24/go-risky-plumbers/pkg/risks"
)

// handle main page
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// handle invalid paths
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the Risky Plumbers Home Page! \nUse /v1/risks for accessing risks or creating new risks!")
}

// main function for risky plumber application
func main() {
	mux := http.NewServeMux()

	// Home Page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rootHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle routes for /v1/risks
	mux.HandleFunc("/v1/risks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			risks.GetRisks(w, r)
		case http.MethodPost:
			risks.CreateRisk(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle routes for /v1/risks/<uuid>
	mux.HandleFunc("/v1/risks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			risks.GetRiskByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Get the port number, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server and listen for requests
	log.Printf("Starting server and listening on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

}
