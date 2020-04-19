package main

import (
	"github.com/dungeonsnd/gocom/media/thumbnail"
	"github.com/dungeonsnd/gocom/test/encrypt/aes"
	"github.com/dungeonsnd/gocom/test/encrypt/rsa"
)

func main() {
	aes.Run()
	rsa.Run()
	thumbnail.ResizeImage("", "", 0)
}
