package padding

import (
	"bytes"
	"errors"
)

func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Pkcs5Unpadding(origData []byte) (dst []byte, err error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if unpadding >= length {
		err = errors.New("unpadding length error")
		return
	}
	dst = origData[:(length - unpadding)]
	return
}
