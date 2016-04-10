package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})

	a_holder := make([]byte, 10)
	my_myreader := MyReader{}
	my_myreader.Read(a_holder)
	fmt.Printf("%q", a_holder[:])
}
