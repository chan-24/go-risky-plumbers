package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Risky Plumbers Home Page! \nUse /v1/risks for accessing risks or creating new risks!")
}

func main() {

    // Get the port number, default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }
	
	http.HandleFunc("/",rootHandler)

	// Start server and listen for requests
	log.Printf("Starting server and listening on port %s", port)
	if err:= http.ListenAndServe(":"+port, nil); err!=nil {
		log.Fatalf("Could not start server: %s", err)
	}

}