package aes

import (
	"fmt"

	"github.com/dungeonsnd/gocom/encrypt/aes"
	"github.com/dungeonsnd/gocom/encrypt/encoding"
)

func Run() bool {
	fmt.Printf("\n-------------------- test encrypt aes --------------------\n")
	plain := []byte("This is a message.")
	key, _ := encoding.HexDecode("6368616e676520746869732070617373")
	iv, _ := encoding.HexDecode("3b9e61ed65ec555f43f9fcb41d5dde3a")

	cipherBytes, err := aes.AesEncrypt(plain, key, iv)
	if err != nil {
		fmt.Printf("AesEncrypt failed, err =%v \n", err)
	} else {
		fmt.Printf("cipherBytes =%v \n", cipherBytes)
	}
	return true
}
