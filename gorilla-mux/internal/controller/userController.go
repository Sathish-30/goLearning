package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sathish-30/gorilla-mux/internal/database"
	"github.com/sathish-30/gorilla-mux/internal/model"
)

func HandleUserHomeRoute(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "In user home route",
	})
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		log.Println("Unable to convert data into user object")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(map[string]string{
			"Error": "Invalid data fields",
		})
		return
	}
	result := database.DB.Create(&admin)
	if result.Error != nil {
		log.Print(result.Error)
	} else {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Added data to Db",
		})
	}

}
