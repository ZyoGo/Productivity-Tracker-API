package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Token  string
	UserID string
}

type Claims struct {
	Username           string
	UserId             string
	jwt.StandardClaims // jwt.StandardClaims is an interface
}
