package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Jwt       string
	ExpiresAt time.Time
}

type jwtData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type AuthData struct {
	IsAuthenticated bool
	Username        string
	Email           string
	Id              int
}
