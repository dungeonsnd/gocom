package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/dungeonsnd/gocom/file/fileutil"
)

func ReadRsaPubKey(priKeyFileName string, pubKeyFileName string) (error, string) {

	publicKey, err := fileutil.ReadFromFile(pubKeyFileName)
	if err != nil {
		err1 := GenRsaKey(priKeyFileName, pubKeyFileName, 2048)
		if err1 != nil {
			return err1, ""
		} else {
			var err2 error
			publicKey, err2 = fileutil.ReadFromFile(pubKeyFileName)
			if err2 != nil {
				return err2, ""
			}

		}
	}
	return nil, string(publicKey)
}

func ReadRsaKeys(priKeyFileName string, pubKeyFileName string) (error, []byte, []byte) {

	priKeyBytes, err := fileutil.ReadFromFile(priKeyFileName)
	if err != nil {
		return err, nil, nil
	}

	publicKey, err := fileutil.ReadFromFile(pubKeyFileName)
	if err != nil {
		return err, nil, nil
	}
	return nil, priKeyBytes, publicKey
}

func GenRsaKey(priKeyFileName string, pubKeyFileName string, bits int) error {
	// gen pri
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return err
	}
	priBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	priKeyBytes := pem.EncodeToMemory(priBlock)
	err = fileutil.WriteToFile(priKeyFileName, priKeyBytes, true)
	if err != nil {
		return err
	}

	// gen pub
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKeyBytes := pem.EncodeToMemory(publicBlock)
	err = fileutil.WriteToFile(pubKeyFileName, pubKeyBytes, true)
	if err != nil {
		return err
	}
	return nil
}

func GenRsaKeyToString(bits int) (string, string, error) {
	// gen pri
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}
	priBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	priKeyBytes := pem.EncodeToMemory(priBlock)

	// gen pub
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKeyBytes := pem.EncodeToMemory(publicBlock)
	return string(priKeyBytes), string(pubKeyBytes), nil
}

func GenRsaKeyPKCS1(priKeyFileName string, pubKeyFileName string, bits int) error {
	// gen pri
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	priKeyBytes := pem.EncodeToMemory(priBlock)
	err = fileutil.WriteToFile(priKeyFileName, priKeyBytes, true)
	if err != nil {
		return err
	}

	// gen pub
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKeyBytes := pem.EncodeToMemory(publicBlock)
	err = fileutil.WriteToFile(pubKeyFileName, pubKeyBytes, true)
	if err != nil {
		return err
	}
	return nil
}
