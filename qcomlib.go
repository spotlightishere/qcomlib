package qcomlib

import "github.com/google/gousb"

type QcomDevice struct {
	in  *gousb.InEndpoint
	out *gousb.OutEndpoint
}

func New(in *gousb.InEndpoint, out *gousb.OutEndpoint) *QcomDevice {
	return &QcomDevice{
		in:  in,
		out: out,
	}
}

func (q *QcomDevice) Write(data []byte) error {
	_, err := q.out.Write(data)
	return err
}

func (q *QcomDevice) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	num, err := q.in.Read(buf)
	if err != nil {
		return nil, err
	}

	// Return only read bytes.
	return buf[:num], nil
}
