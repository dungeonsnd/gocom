package typeconv

import (
	"bytes"
	"encoding/binary"
)

func Int8ToBytes(x int8) (buf []byte, err error) {
	b_buf := new(bytes.Buffer)
	err = binary.Write(b_buf, binary.BigEndian, x)
	buf = b_buf.Bytes()
	return
}

func BytesToInt8(buf []byte) (x int8, err error) {
	b_buf := bytes.NewBuffer(buf)
	err = binary.Read(b_buf, binary.BigEndian, &x)
	return
}

func Int16ToBytes(x int16) (buf []byte, err error) {
	b_buf := new(bytes.Buffer)
	err = binary.Write(b_buf, binary.BigEndian, x)
	buf = b_buf.Bytes()
	return
}

func BytesToInt16(buf []byte) (x int16, err error) {
	b_buf := bytes.NewBuffer(buf)
	err = binary.Read(b_buf, binary.BigEndian, &x)
	return
}

func Int32ToBytes(x int32) (buf []byte, err error) {
	b_buf := new(bytes.Buffer)
	err = binary.Write(b_buf, binary.BigEndian, x)
	buf = b_buf.Bytes()
	return
}

func BytesToInt32(buf []byte) (x int32, err error) {
	b_buf := bytes.NewBuffer(buf)
	err = binary.Read(b_buf, binary.BigEndian, &x)
	return
}
