package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sathish-30/go-testing/internal/handler"
)

func main() {
	http.HandleFunc("/user", handler.UserHomeRoute)
	port := fmt.Sprintf(":%d", 8000)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
