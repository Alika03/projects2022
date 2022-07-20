package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func CreateKey(privateFileName, publicFileName string) error {
	// create rsa key
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}
	// get public key from private key
	publicKey := &privateKey.PublicKey

	// dump private key to file
	var privateKeyBytes = x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err = writeToFile(privateFileName, privateKeyBlock); err != nil {
		return err
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	if err = writeToFile(publicFileName, publicKeyBlock); err != nil {
		return err
	}

	return nil
}

func writeToFile(fileName string, keyBlock *pem.Block) error {
	keyFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err = pem.Encode(keyFile, keyBlock); err != nil {
		return err
	}

	return nil
}
