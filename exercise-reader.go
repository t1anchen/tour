package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 'A'
	}
	return len(buf), nil
}

func main() {
	reader.Validate(MyReader{})
}
