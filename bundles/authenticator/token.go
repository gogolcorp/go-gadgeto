package authenticator

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// Claim is the struct for the jwt claim
type Claim struct {
	Email string
	jwt.StandardClaims
}

// GenerateToken creates a JWT with email and expiration time in the payload
func GenerateToken(email string) (string, error) {
	err := initRsaKeys()
	if err != nil {
		return "", errors.New("couldn't init rsa keys")
	}

	validTime, _ := strconv.ParseInt(goDotEnvVariable("TOKEN_VALID_DURATION"), 10, 64)
	// Generate Expiration date
	expirationTime := time.Now().Add(time.Duration(validTime) * time.Minute)

	claims := &Claim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// JWT takes unix timestamps
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(rsa.PrivateKey)

	if err != nil {
		log.Error("error while generating token: ", err)
		return "", err
	}

	return tokenString, nil
}

// DecodeToken decode and validates a token
func DecodeToken(tokenString string) (*jwt.Token, *Claim, error) {
	err := initRsaKeys()
	if err != nil {
		return nil, &Claim{}, errors.New("couldn't init rsa keys")
	}

	claims := &Claim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return rsa.PublicKey, nil
	})
	if err != nil {
		log.Error("Couldn't parse the token : ", err)
		return nil, &Claim{}, err
	}

	return token, claims, nil
}
