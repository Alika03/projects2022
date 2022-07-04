package utils

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go/v4"
)

type Jwt struct {
	Key    *rsa.PrivateKey
	Claims *jwt.StandardClaims
}

func NewJwt(key *rsa.PrivateKey, claims *jwt.StandardClaims) *Jwt {
	return &Jwt{
		Key:    key,
		Claims: claims,
	}
}

func (j *Jwt) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, j.Claims)

	return token.SignedString(j.Key)
}

func (j *Jwt) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return j.Key, nil })
}
