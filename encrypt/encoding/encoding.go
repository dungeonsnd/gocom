package encoding

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

func HexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

func HexDecode(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

func Base32Encode(data string) string {
	return base32.StdEncoding.EncodeToString([]byte(data))
}

func Base32Decode(str string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(str)
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func JsonEncode(m interface{}) (buf []byte, err error) {
	return json.Marshal(m)
}

func JsonDecode(buf []byte) (interface{}, error) {
	var m interface{}
	err := json.Unmarshal(buf, &m)
	return m, err
}

func JsonHexEncode(structObj interface{}) (error, string) {
	b, err := json.Marshal(structObj)
	if err != nil {
		return err, ""
	}
	return nil, hex.EncodeToString(b)
}
