package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	app := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	app.Handle("/static/", http.StripPrefix("/static/", fs))
	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "In Home route",
		})
	})
	http.ListenAndServe(":8000", app)
}
