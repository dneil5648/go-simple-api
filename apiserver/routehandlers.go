package apiserver

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandleHealthRoute(w http.ResponseWriter, r *http.Request) {

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
