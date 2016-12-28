package main

import "fmt"

func main() {
	//一个nil值的slice的行为和其它任意0长度的slice一样

	s := []int{}
	//s2 := s[:5]
	fmt.Println(cap(s))
	s2 := append(s, 7)
	fmt.Println(s2)

	var t []int
	t2 := append(t, 6)
	fmt.Println(t2)
}
