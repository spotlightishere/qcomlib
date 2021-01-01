package qcomlib

import (
	"bytes"
	"errors"
)

var (
	RequestHello = []byte{0x00, 0x00, 0x10, 0x00, 0x01, 0xff, 0x00, 0x00, 0xa8, 0x18, 0xda, 0x8d, 0x6c, 0x02, 0x00, 0x00}
)

var (
	ResponseHello = []byte{0x07, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x78, 0x56, 0x34, 0x12}
)

var (
	ErrInvalidResponse = errors.New("device did not respond with expected data")
)

func (q *QcomDevice) Hello() error {
	if err := q.Write(RequestHello); err != nil {
		return err
	}

	response, err := q.Read()
	if err != nil {
		return err
	}

	if bytes.Compare(response, ResponseHello) == 0 {
		return nil
	} else {
		return ErrInvalidResponse
	}
}
