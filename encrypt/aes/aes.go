package aes

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/dungeonsnd/gocom/encrypt/encoding"
	pkcs "github.com/dungeonsnd/gocom/encrypt/padding"
)

func AesDecryptForB64ByPkcs5Padding(origDataB64 string, key []byte, iv []byte) ([]byte, error) {

	p, err := encoding.Base64Decode(origDataB64)
	if err != nil {
		return nil, err
	}

	return AesDecryptByPkcs5Padding(p, key, iv)
}

func AesDecryptForB64ByPkcs7Padding(origDataB64 string, key []byte, iv []byte) ([]byte, error) {

	p, err := encoding.Base64Decode(origDataB64)
	if err != nil {
		return nil, err
	}

	return AesDecryptByPkcs7Padding(p, key, iv)
}

func AesDecryptForB64AndHexKeyByPkcs5Padding(origDataB64 string, keyHex string, ivHex string) ([]byte, error) {
	k, err := encoding.HexDecode(keyHex)
	if err != nil {
		return nil, err
	}
	i, err := encoding.HexDecode(ivHex)
	if err != nil {
		return nil, err
	}

	p, err := encoding.Base64Decode(origDataB64)
	if err != nil {
		return nil, err
	}

	return AesDecryptByPkcs5Padding(p, k, i)
}

func AesDecryptForB64AndHexKeyByPkcsP7adding(origDataB64 string, keyHex string, ivHex string) ([]byte, error) {
	k, err := encoding.HexDecode(keyHex)
	if err != nil {
		return nil, err
	}
	i, err := encoding.HexDecode(ivHex)
	if err != nil {
		return nil, err
	}

	p, err := encoding.Base64Decode(origDataB64)
	if err != nil {
		return nil, err
	}

	return AesDecryptByPkcs7Padding(p, k, i)
}

func AesEncryptForHexKeyByPkcs5Padding(origData []byte, keyHex string, ivHex string) ([]byte, error) {
	k, err := encoding.HexDecode(keyHex)
	if err != nil {
		return nil, err
	}
	i, err := encoding.HexDecode(ivHex)
	if err != nil {
		return nil, err
	}
	return AesEncryptByPkcs5Padding(origData, k, i)
}

func AesEncryptForHexKeyByPkcs7Padding(origData []byte, keyHex string, ivHex string) ([]byte, error) {
	k, err := encoding.HexDecode(keyHex)
	if err != nil {
		return nil, err
	}
	i, err := encoding.HexDecode(ivHex)
	if err != nil {
		return nil, err
	}
	return AesEncryptByPkcs7Padding(origData, k, i)
}

func AesEncryptByPkcs5Padding(origData []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs.Pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesEncryptByPkcs7Padding(origData []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs.Pkcs7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecryptByPkcs5Padding(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	// fmt.Printf("AesDecrypt, crypted=%v, key=%v, iv=%v\n", string(crypted), string(key), string(iv))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = pkcs.Pkcs5Unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

func AesDecryptByPkcs7Padding(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	// fmt.Printf("AesDecrypt, crypted=%v, key=%v, iv=%v\n", string(crypted), string(key), string(iv))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = pkcs.Pkcs7Unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}
