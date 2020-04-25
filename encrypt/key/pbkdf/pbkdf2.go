package pbkdf

import (
	"hash"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(pwd []byte, salt []byte, rounds int, resultLen int, h func() hash.Hash) []byte {
	key := pbkdf2.Key([]byte(pwd), salt, rounds, resultLen, h)
	return key
}
