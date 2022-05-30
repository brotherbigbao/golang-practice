package main

import (
	"fmt"
	"strings"
)

func main() {
	example()
}

//中文空格也可以去除
func example() {
	s := " \t\n a lone gopher码农　 \n\t\r\n"
	fmt.Println(s)

	s = strings.TrimSpace(s)
	if s == "a lone gopher码农" {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	fmt.Print(s)
}
