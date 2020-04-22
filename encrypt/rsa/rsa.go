package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"

	gcsha256 "github.com/dungeonsnd/gocom/encrypt/hash/sha256"
)

func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key decode error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	// fmt.Printf("RsaEncrypt, pub.Size()-11=%v\n", pub.Size()-11)

	// ecrypt by pkcs1 padding
	b, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	if err != nil {
		fmt.Printf("RsaEncrypt, ParsePKIXPublicKey err=%v\n", err)
		return nil, err
	}
	return b, err
}

func RsaDecrypt(ciphertext []byte, priKeyBytes []byte) ([]byte, error) {

	block, _ := pem.Decode(priKeyBytes)
	if block == nil {
		return nil, errors.New("private key decode error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	b, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return nil, err
	}

	return b, err
}

func RsaDecryptByDefaultKey(b64CipherText string, sumHext string) ([]byte, error) {

	priDefault := "./pri.key"
	pubDefault := "./pub.key"
	err, priKeyBytes, _ := ReadRsaKeys(priDefault, pubDefault)
	if err != nil {
		err := GenRsaKey(priDefault, pubDefault, 2048)
		if err != nil {
			return nil, err
		}
	}
	dd, err := base64.StdEncoding.DecodeString(b64CipherText)
	if err != nil {
		return nil, err
	}

	origData, err := RsaDecrypt(dd, priKeyBytes)
	if err != nil {
		return nil, err
	}
	sumData := gcsha256.HashHex(origData, 1)
	if sumData != sumHext {
		return nil, errors.New("hash check failed")
	}

	return origData, nil
}
