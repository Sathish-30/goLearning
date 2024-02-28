package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sathish-30/gorilla-mux/internal/database"
	"github.com/sathish-30/gorilla-mux/internal/routers"
)

func init() {
	database.ConnectToDb()
	database.AutoMigrate()
	log.Println("Database connected and table migrated")
}

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	routers.UserRoutes(router)
	http.ListenAndServe(":8080", router)
}
