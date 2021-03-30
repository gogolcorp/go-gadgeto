package authenticator

import (
	"github.com/dgrijalva/jwt-go"
)
type authenticatorInterface interface {
	GenerateToken(string) (error, string)
	DecodeToken(string)
	HashPassword(string, string) error
}

// Claim is the struct for the jwt claim
type Claim struct {
	Email string
	jwt.StandardClaims
}

// Authenticator is the struct for the authenticator bundle
type Authenticator struct {}

// New returns a new authenticator value
func New() Authenticator {
	auth := Authenticator{}
	return auth
}