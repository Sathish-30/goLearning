package auth

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Name string
	jwt.RegisteredClaims
}
