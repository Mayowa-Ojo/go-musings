package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var asciiUppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var asciiLowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var asciiUppercaseLen = len(asciiUppercase)
var asciiLowercaseLen = len(asciiLowercase)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func rot13(b byte) byte {
	pos := bytes.IndexByte(asciiUppercase, b)
	if pos != -1 {
		return asciiUppercase[(pos+13)%asciiUppercaseLen]
	}
	pos = bytes.IndexByte(asciiLowercase, b)
	if pos != -1 {
		return asciiLowercase[(pos+13)%asciiLowercaseLen]
	}
	return b
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}
