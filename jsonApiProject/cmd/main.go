package main

import (
	"log"
	"net/http"

	"github.com/sathish-30/jsonApiProject/internal/handler"
	"github.com/sathish-30/jsonApiProject/internal/util"
)

type Server struct {
	listenAddr string
}

type ApiError struct {
	Error string
}

func main() {
	server := getServer(":8000")
	server.run()
}

func getServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
	}
}

func makeHttpHandleFunc(f handler.ApiFuncCall) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			util.WriteJson(w, http.StatusBadRequest, ApiError{
				Error: err.Error(),
			})
		}
	}
}

func (server *Server) run() {
	router := http.NewServeMux()
	router.HandleFunc("/", makeHttpHandleFunc(handler.HomeRouteHandler))
	log.Println("Json api server is running on the port ", server.listenAddr)
	http.ListenAndServe(server.listenAddr, router)
}
