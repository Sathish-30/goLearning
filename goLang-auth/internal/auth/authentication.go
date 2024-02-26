package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
)

func AuthWithJWT(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		log.Fatal(cookie)
		if err != nil {
			if err != http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value
		token, err := jwt.ParseWithClaims(tokenString, , func(token *jwt.Token) (interface{}, error) {
			secretKey := []byte(os.Getenv("SECRET"))
			return secretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
