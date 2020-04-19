package rsa

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/dungeonsnd/gocom/encrypt/rsa"
)

func Run() bool {
	fmt.Printf("\n-------------------- test encrypt rsa --------------------\n")

	err, priKeyBytes, publicKey := rsa.ReadRsaKeys("./pri.key", "./pub.key")
	if err != nil {
		err := rsa.GenRsaKey("./pri.key", "./pub.key", 2048)
		if err != nil {
			fmt.Printf("EncData, genRsaKey err=%v\n", err)
			return false
		} else {
			fmt.Printf("EncData, genRsaKey sucess\n")
		}
	}
	err, priKeyBytes, publicKey = rsa.ReadRsaKeys("./pri.key", "./pub.key")
	if err != nil {
		fmt.Printf("EncData, readRsaKey err=%v\n", err)
		return false
	}

	// d := []byte("12345679012345679012345679012345679012345679012345679012345679012345679012345679012345679012345679012345679012a234567")
	d := []byte("123abc...")
	sum := sha256.Sum256(d)
	fmt.Printf("Plain SHA256: %s\n", hex.EncodeToString(sum[:]))

	data, err := rsa.RsaEncrypt(d, publicKey)
	if err != nil {
		fmt.Printf("EncData, RsaEncrypt err=%v\n", err)
		return false
	}
	fmt.Printf("Encrypted(%v): %s\n", len(data), base64.StdEncoding.EncodeToString(data))

	origData, err := rsa.RsaDecrypt(data, priKeyBytes)
	if err != nil {
		fmt.Printf("EncData, RsaDecrypt err=%v\n", err)
		return false
	}
	sum = sha256.Sum256(origData)
	fmt.Printf("Decrypted: %s\n", string(origData))
	fmt.Printf("Decrypted SHA256: %s\n", hex.EncodeToString(sum[:]))

	// fmt.Printf("------------- Decrypt base64 string From JS ------------\n")

	// h := "jLdaKFr5XMCQq+kEZdzNQeaSWPPXZAK738VPLMv6ibp18AISU2OFOETk47VYC5ch2grWW4JQ9vlu0VJErp04J4yw4wM2JABS0ukieWST5JAfXZuHFyH2KP8Ys4qmN7YCubzZ4+ozhdI6zmApSsHWTvUdLgCrVHk9OWOpZgbwIyM="
	// dd, err := base64.StdEncoding.DecodeString(h)
	// if err != nil {
	// 	fmt.Printf("EncData, base64.StdEncoding.DecodeString err=%v\n", err)
	// 	return false
	// }

	// od, err := rsa.RsaDecrypt(dd, priKeyBytes)
	// if err != nil {
	// 	fmt.Printf("EncData, RsaDecrypt err=%v\n", err)
	// 	return false
	// }
	// fmt.Printf("Decrypted text: %s\n", string(od))

	return true
}
