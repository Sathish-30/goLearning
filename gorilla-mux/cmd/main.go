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
	routers.UserRoutes(router)
	http.ListenAndServe(":8000", router)
}
