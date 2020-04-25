package aes

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/dungeonsnd/gocom/encrypt/encoding"
	pkcs7 "github.com/dungeonsnd/gocom/encrypt/padding"
)

func AesDecryptForB64(origDataB64 string, key []byte, iv []byte) ([]byte, error) {

	p, err := encoding.Base64Decode(origDataB64)
	if err != nil {
		return nil, err
	}

	return AesDecrypt(p, key, iv)
}

func AesEncryptForHexKey(origData []byte, keyHex string, ivHex string) ([]byte, error) {
	k, err := encoding.HexDecode(keyHex)
	if err != nil {
		return nil, err
	}
	i, err := encoding.HexDecode(ivHex)
	if err != nil {
		return nil, err
	}
	return AesEncrypt(origData, k, i)
}

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
