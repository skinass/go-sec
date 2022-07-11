package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
	var b *bytes.Buffer = bytes.NewBuffer(data)
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}
	res, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	r.Close()
	return res, nil
}

func main() {
	data := bytes.Repeat([]byte("0"), 1024*1024)
	fmt.Printf("data len = %d bytes\n", len(data))

	c, _ := compress(data)
	fmt.Printf("compressed data len = %d bytes\n", len(c))
	fmt.Printf("compressed data = %+v\n", c)
}
