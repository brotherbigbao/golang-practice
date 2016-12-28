package main

import "fmt"

func main() {
	s := "left"
	t := s
	s += " right"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(string(s[1]))
	fmt.Println(t)
}