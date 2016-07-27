package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13r rot13Reader) Read(buf []byte) (int, error) {
	n_read, err := r13r.r.Read(buf)
	for i, c := range buf {
		if c >= 'A' && c <= 'Z' {
			buf[i] = r13r.Rot13(c, 'A')
		}
		if c >= 'a' && c <= 'z' {
			buf[i] = r13r.Rot13(c, 'a')
		}
	}
	return n_read, err
}

func (r13r rot13Reader) Rot13(src byte, base byte) byte {
	return (src-base+13)%26 + base
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// $ go run exercise-rot-reader.go
// You cracked the code!
