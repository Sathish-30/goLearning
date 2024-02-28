package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sathish-30/gorilla-mux/internal/model"
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

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Unable to encode request body",
		})
		return
	}
	log.Println(user)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Unable to encode request body",
		})
		return
	}
}
