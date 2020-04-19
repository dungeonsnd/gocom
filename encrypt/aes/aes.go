package aes

import (
	"crypto/aes"
	"crypto/cipher"

	pkcs7 "github.com/dungeonsnd/gocom/encrypt/padding"
)

func AesEncrypt(origData []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs7.Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
func AesDecrypt(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	// fmt.Printf("AesDecrypt, crypted=%v, key=%v, iv=%v\n", string(crypted), string(key), string(iv))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = pkcs7.Unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}
