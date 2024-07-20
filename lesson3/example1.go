package main

import (
	"io"
	"os"
	"strings"

	// подключение стороннего пакета
	"github.com/jreisinger/gokatas/rot13"
)

func main() {
	// работа со сторонним пакетом
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13.Reader{R: s}

	io.Copy(os.Stdout, &r)
}
