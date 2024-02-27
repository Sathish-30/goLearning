package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sathish-30/gorilla-mux/internal/routers"
)

func main() {
	router := mux.NewRouter()
	routers.UserRoutes(router)
	http.ListenAndServe(":8000", router)
}
