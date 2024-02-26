package auth

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Name string
	jwt.StandardClaims
}
