package routers

import (
	"github.com/gorilla/mux"
	"github.com/sathish-30/gorilla-mux/internal/controller"
)

func UserRoutes(router *mux.Router) {
	userRoute := router.PathPrefix("/user").Subrouter()
	// User home route
	userRoute.HandleFunc("/", controller.HandleUserHomeRoute).Methods("GET")

	userRoute.HandleFunc("/login", controller.HandleUserLogin).Methods("POST")
}
