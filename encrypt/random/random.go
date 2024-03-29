package random

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	math_rand "math/rand"
	"time"

	"github.com/dungeonsnd/gocom/encrypt/hash/sha256"
	uuid "github.com/satori/go.uuid"
)

func Random(byteLen uint16) []byte {

	b := make([]byte, byteLen)
	n, err := rand.Read(b) //在byte切片中随机写入元素
	if err != nil {
		return nil
	}
	if n != int(byteLen) {
		return nil
	}
	return b
}

func RandomHex(byteLen uint16) string {

	b := make([]byte, byteLen)
	n, err := rand.Read(b) //在byte切片中随机写入元素
	if err != nil {
		return ""
	}
	if n != int(byteLen) {
		return ""
	}
	return hex.EncodeToString(b)
}

func RandomB64(byteLen uint16) string {

	b := make([]byte, byteLen)
	n, err := rand.Read(b) //在byte切片中随机写入元素
	if err != nil {
		return ""
	}
	if n != int(byteLen) {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}

func RandMixByte() []byte {
	u1 := uuid.NewV4()
	u2 := uuid.NewV1()
	u := append(u1[0:], u2[0:]...)

	u3 := make([]byte, 32)
	_, err := rand.Read(u3)
	if err == nil {
		u = append(u, u3...)
	}
	return u
}

func RandMixHashHex() string {
	u1 := uuid.NewV4()
	u2 := uuid.NewV1()
	u := append(u1[0:], u2[0:]...)

	u3 := make([]byte, 32)
	_, err := rand.Read(u3)
	if err == nil {
		u = append(u, u3...)
	}
	return sha256.HashHex(u, 1)
}

func RandNum(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GenUUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}
