package authenticator

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// Rsa is the struct to get the rsa keys used to generate and verify tokens from the environment variables
type Rsa struct {
	PublicKeyPath  string
	PrivateKeyPath string
	PublicKey      interface{}
	PrivateKey     interface{}
}

var rsa Rsa

func initRsaKeys() error {
	if rsa.PrivateKey != nil && rsa.PublicKey != nil {
		return nil
	}

	rsa.PublicKeyPath = goDotEnvVariable("RSA_PUBLIC_PATH")
	rsa.PrivateKeyPath = goDotEnvVariable("RSA_PRIVATE_PATH")

	if err := parsePublicKey(&rsa.PublicKey); err != nil {
		log.Error("Couldnt parse public.pem", err)
		return err
	}

	if err := parsePrivateKey(&rsa.PrivateKey); err != nil {
		log.Error("Couldn't parse private.pem", err)
		return err
	}

	return nil
}

func getPublicPemFile(fileData *[]byte) error {
	var err error
	*fileData, err = ioutil.ReadFile(rsa.PublicKeyPath)

	return err
}

func parsePublicKey(publicKey *interface{}) error {
	var fileData []byte
	if err := getPublicPemFile(&fileData); err != nil {
		log.Error("Couldnt get public.pem file")
		return err
	}
	var err error
	*publicKey, err = jwt.ParseRSAPublicKeyFromPEM(fileData)

	return err
}

func parsePrivateKey(privateKey *interface{}) error {
	var fileData []byte
	if err := getPrivatePemFile(&fileData); err != nil {
		log.Error("Couldn't get private.pem file")
		return err
	}

	var err error
	*privateKey, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword(fileData, goDotEnvVariable("RSA_PASSWORD"))

	return err
}

func getPrivatePemFile(fileData *[]byte) error {
	var err error
	*fileData, err = ioutil.ReadFile(rsa.PrivateKeyPath)

	return err
}
