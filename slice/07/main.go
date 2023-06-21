package main

import "fmt"

func main() {
	s := make([]int, 0)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	if s == nil {
		fmt.Println("s is nil.")
	} else {
		fmt.Println("s is not nil.")
	}

	s = append(s, 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	if s == nil {
		fmt.Println("s is nil.")
	} else {
		fmt.Println("s is not nil.")
	}
}
