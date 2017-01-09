package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T\n", w)
	rw := w.(io.ReadWriter)
	fmt.Printf("%T\n", rw)
	rw = io.ReadWriter(rw)
	fmt.Printf("%T\n", rw)

	w = new(ByteCounter)
	fmt.Printf("%T\n", w)
	rw, ok := w.(io.ReadWriter)
	fmt.Printf("%T,%v\n", rw, ok)
}

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}