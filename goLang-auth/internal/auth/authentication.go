package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string = string(os.Getenv("SECRET"))

func AuthWithJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access: Token missing\n")
			return
		}

		// Parse the JWT token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access: Invalid token\n")
			return
		} else if claims, ok := token.Claims.(*CustomClaims); ok {
			log.Println(claims.Name, claims.RegisteredClaims)
		} else {
			log.Fatal("Unknown claim types")
		}
		next.ServeHTTP(w, r)
	})
}
