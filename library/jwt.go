package library

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//Claim struct
type Claim struct {
	Issuer   string
	Audience string
	Subject  string
	Expired  time.Duration
	Lock     *sync.RWMutex
}

//AccessTokenResponse struct
type AccessTokenResponse struct {
	Error       error
	AccessToken string
}

//InitKeys function
func InitKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	filepath, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privateKeyPath := filepath + "/secret/app.rsa"
	publicKeyPath := filepath + "/secret/app.rsa.pub"

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		log.Fatal(err)
	}

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)

	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	if err != nil {
		log.Fatal(err)
	}

	return signKey, verifyKey
}

//NewClaim function
func NewClaim(issuer string, audience string, subject string, expired time.Duration) *Claim {
	return &Claim{
		Issuer:   issuer,
		Audience: audience,
		Subject:  subject,
		Expired:  expired,
		Lock:     new(sync.RWMutex),
	}
}

//GenerateToken function
func (cl *Claim) GenerateToken(signKey *rsa.PrivateKey) <-chan AccessTokenResponse {
	result := make(chan AccessTokenResponse)
	go func() {
		cl.Lock.Lock()
		defer close(result)
		defer cl.Lock.Unlock()
		token := jwt.New(jwt.SigningMethodRS256)
		claims := make(jwt.MapClaims)
		claims["iss"] = cl.Issuer
		claims["aud"] = cl.Audience
		claims["exp"] = time.Now().Add(cl.Expired).Unix()
		claims["iat"] = time.Now().Unix()
		claims["sub"] = cl.Subject
		token.Claims = claims

		tokenString, err := token.SignedString(signKey)
		if err != nil {
			result <- AccessTokenResponse{Error: err, AccessToken: ""}
			return
		}

		result <- AccessTokenResponse{Error: nil, AccessToken: tokenString}
	}()
	return result
}
