package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
)

func main() {
	// 断言具体类型
	var w io.Writer
	w = os.Stdout
	f, ok1 := w.(*os.File)
	c, ok2 := w.(*bytes.Buffer)
	fmt.Println(w==f)
	fmt.Println(f, ok1)
	fmt.Println(c, ok2)
}