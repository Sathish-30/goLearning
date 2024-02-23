package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sathish-30/go-testing/internal/handler"
)

func loggerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("The Url Path is %s and the method is %s", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}

func main() {
	app := http.NewServeMux()
	// Logger middleware
	wrappedServer := loggerMiddleWare(app)

	app.HandleFunc("/user/", handler.HandleGetUserById)
	app.HandleFunc("/user", handler.UserHomeRoute)
	app.HandleFunc("POST /addUser", handler.AddUserHandler)
	app.HandleFunc("DELETE /deleteUser/", handler.DeleteUserByIdHandler)


	port := fmt.Sprintf(":%d", 8000)
	if err := http.ListenAndServe(port, wrappedServer); err != nil {
		log.Fatal(err)
	}
   
}
