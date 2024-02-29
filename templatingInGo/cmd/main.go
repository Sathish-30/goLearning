package main

import (
	"net/http"
	"path"
	"text/template"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", showBooks)
	http.ListenAndServe(":8000", server)
}

func showBooks(w http.ResponseWriter, r *http.Request) {
	book := struct {
		BookName   string
		AuthorName string
	}{
		BookName:   "Alchemist",
		AuthorName: "Sathish",
	}
	// Where the path gets defined from the starting relative path
	fp := path.Join("internal", "templates", "index.html")
	template, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := template.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
