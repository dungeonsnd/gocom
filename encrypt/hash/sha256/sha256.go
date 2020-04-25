package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Hash(src []byte, rounds int) []byte {
	dst := src
	for i := 0; i < rounds; i++ {
		t := sha256.Sum256(dst)
		dst = t[0:]
	}
	return dst
}

func HashHex(src []byte, rounds int) string {
	h := Hash(src, rounds)
	return hex.EncodeToString(h)
}

func CheckHash(dataByte []byte, expectedHashHex string, minExpectedHashHexLen uint) (bool, string) {
	dataHash := HashHex(dataByte, 1)
	if len(expectedHashHex) < minExpectedHashHexLen || strings.Index(dataHash, expectedHashHex) != 0 {
		return false, dataHash
	} else {
		return true, dataHash
	}
}
