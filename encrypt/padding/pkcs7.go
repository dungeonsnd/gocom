// This package implements PKCS#7 padding to byte slices. Padding works by
// calculating the amount of needed padding and repeating that number to fill in
// the rest of the slice up to the block size. This way, in order to unpad the
// slice, you check the value held in the last slot and then remove that many
// bytes from the end.
//
// By defniition, PKCS#7 only padds for block sizes between 1 and 255 inclusive.
// If the supplied byte slice is a multiple of the block size, N, an extra N
// amount of bytes is appended all of value N.
//
// Please review the tests for this package for examples.
package pkcs7

import (
	"bytes"
	"errors"
)

// Pad takes a source byte slice and a block size. It will determine the needed
// amount of padding, n, and appends byte(n) to the source n times.
//
// Example Input: Block Size 8, Source {0xDE, 0xAD, 0xBE, 0xEF}
//
// Expected Output: {0xDE, 0xAD, 0xBE, 0xEF, 0x04, 0x04, 0x04, 0x04}
//
func Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// Unpad takes a source byte slice and will remove any padding added according
// to PKCS#7 specifications. An error is returned for invalid padding.
func Unpadding(src []byte) (dst []byte, err error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding >= length {
		err = errors.New("unpadding length error!")
		return
	}
	dst = src[:(length - unpadding)]
	return
}
