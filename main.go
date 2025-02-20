package main

import (
	"encoding/json"
	"go-simple-api/apiserver"
	"net/http"
	"time"
)

func handleHealthRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	healthStatus := map[string]string{
		"status": "healthy",
		"time":   time.Now().UTC().String(),
	}

	//Set Headers
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	//Encode and write Response
	json.NewEncoder(w).Encode(healthStatus)
}

func main() {
	server := apiserver.CreateNewServer(":8080")

	server.AddRoute("/health", handleHealthRoute)

	server.Run()
}
