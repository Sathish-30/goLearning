package util

import (
	"encoding/json"
	"net/http"

	"github.com/sathish-30/jsonApiProject/internal/model"
)

func WriteJson(w http.ResponseWriter, status int, payload any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}

func ReadJson[T model.Book](r *http.Request) (T, error) {
	var retrievedData T
	err := json.NewDecoder(r.Body).Decode(&retrievedData)
	return retrievedData, err
}
