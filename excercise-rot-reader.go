package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M') {
		b += 13
	} else if (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z') {
		b -= 13
	}
	return b
}

func (a rot13Reader) Read(b []byte) (int, error) {
	n, e := a.r.Read(b)
	for i, v := range b {
		b[i] = rot13(v)
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
