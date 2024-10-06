package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("hello world")
	hashAndBroadcast(NewHashReader(payload))
}

type HashReader interface {
	io.Reader
	hash() string
}
type hashReader struct {
	reader *bytes.Reader
	buf    *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {

	return &hashReader{
		reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}

}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(hr HashReader) error {
	hash := hr.hash()

	fmt.Println(hash)

	return broadcast(hr)

}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	fmt.Println("string of bytes", string(b))

	return nil

}
