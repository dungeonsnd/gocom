package io

import (
	"io"
	"net"
	"time"
)

func SetReadTimeout(conn net.Conn, timeoutSec int) {
	conn.SetDeadline(time.Now().Add(time.Duration(timeoutSec) * time.Second))
}

func ReadData(conn net.Conn, total int) ([]byte, error) {
	buf := make([]byte, total)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func WriteData(conn net.Conn, buf []byte) error {
	total := len(buf)
	for nw := 0; nw < total; {
		n, err := conn.Write(buf[nw:])
		if err != nil {
			return err
		}
		nw += n
	}
	return nil
}
