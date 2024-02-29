package util

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, payload any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}
