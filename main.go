package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Allows chaining readers
type alphaReader struct {
	src io.Reader
}

func NewAlphaReader(source io.Reader) *alphaReader {
	return &alphaReader{source}
}

func (a *alphaReader) Read(p []byte) (int, error) {
	// Not empty object
	if len(p) == 0 {
		return 0, nil
	}

	count, err := a.src.Read(p) // Only here string reader stream is called
	if err != nil {
		return count, err
	}

	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') ||
			(p[i] >= 'a' && p[i] <= 'z') {
			continue
		} else {
			p[i] = 0
		}
	}

	return count, io.EOF
}

func main() {
	fmt.Println("Starting program...")

	str := strings.NewReader("This is! quite a nice string")
	fmt.Printf("%T\n", str)
	fmt.Printf("%+v\n", str)
	fmt.Println(str)

	alpha := NewAlphaReader(str)
	io.Copy(os.Stdout, alpha)
	fmt.Println()

	// str := alphareader("This is! quite a nice string")
	// io.Copy(os.Stdout, &str)
	// fmt.Println()

	fmt.Println("Program finished")
}
