package encoding

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/dungeonsnd/gocom/file/fileutil"
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

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func JsonEncode(m interface{}) (buf []byte, err error) {
	return json.Marshal(m)
}

func JsonDecode(buf []byte, m interface{}) error {
	err := json.Unmarshal(buf, &m)
	return err
}

func JsonHexEncode(structObj interface{}) (error, string) {
	b, err := json.Marshal(structObj)
	if err != nil {
		return err, ""
	}
	return nil, hex.EncodeToString(b)
}

func WriteToFileAsJson(filename string, v interface{}, indent string, truncateIfExist bool) error {

	buf, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	err = fileutil.WriteToFile(filename, buf, true)
	if err != nil {
		return err
	}
	return nil
}

func ReadFileJsonToObject(filename string, obj interface{}) error {

	err, buf := fileutil.ReadFromFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return err
	}
	return nil
}
