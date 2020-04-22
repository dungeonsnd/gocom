package hmac

import (
	"crypto/hmac"
	"errors"

	stdsha256 "crypto/sha256"
)

func HmacSha256(msg, secret []byte) ([]byte, error) {
	mac := hmac.New(stdsha256.New, secret)
	n, err := mac.Write(msg)
	if err != nil {
		return nil, err
	}
	if n != len(msg) {
		return nil, errors.New("Written length error.")
	}
	return mac.Sum(nil), nil
}
