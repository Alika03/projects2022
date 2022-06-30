package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"sync"
)

const (
	privateKeyNameOfFile = "privateKey.pem"
	publicKeyNameOfFile  = "publicKey.pem"
)

var keys *Keys

var onceKeys sync.Once

type Keys struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

func GetKeys() *Keys {
	onceKeys.Do(func() {
		keys = &Keys{}
		privateKeyBytes, err := ioutil.ReadFile(privateKeyNameOfFile)
		if err != nil {
			log.Fatalln(err)
		}

		publicKeyBytes, err := ioutil.ReadFile(publicKeyNameOfFile)
		if err != nil {
			log.Fatalln(err)
		}

		// write key to struct
		privateKeyBlock, _ := pem.Decode(privateKeyBytes)
		keys.Private, _ = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)

		publicKeyBlock, _ := pem.Decode(publicKeyBytes)
		publicInterface, _ := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
		keys.Public = publicInterface.(*rsa.PublicKey)

	})

	return keys
}
