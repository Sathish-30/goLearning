package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("GET /user/{id}", handleGetUserById)
	http.ListenAndServe(":8000", server)
}

func handleGetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println(id)
	json.NewEncoder(w).Encode(map[string]any{
		"UserId": id,
	})
}
