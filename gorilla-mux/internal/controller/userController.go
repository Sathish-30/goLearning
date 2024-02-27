package controller

import (
	"encoding/json"
	"net/http"
)

func HandleUserHomeRoute(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "In user home route",
	})
}

func HandleHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "In Home route",
	})
}
