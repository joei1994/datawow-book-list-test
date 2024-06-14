package models

import "github.com/golang-jwt/jwt/v5"

const (
	ExpireCount        = 2
	ExpireRefreshCount = 168
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}
