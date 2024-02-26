package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/sathish-30/golang-auth/initializers"
	"github.com/sathish-30/golang-auth/internal/controller"
	"github.com/sathish-30/golang-auth/internal/database"
)

func init() {
	initializers.LoadEnvVaribales()
	database.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	server := http.NewServeMux()
	// Parsing port number from dotenv file
	portNum, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	port := fmt.Sprintf(":%d", portNum)

	server.HandleFunc("POST /signUp", controller.SignUp)

	server.HandleFunc("POST /login", controller.LoginUser)

	// Listening to the port number which has been obtained from the .env file
	if err := http.ListenAndServe(port, server); err != nil {
		log.Fatal(err)
	}
}
