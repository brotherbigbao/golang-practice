package _2

import "fmt"

func main2() {
	var s []int
	fmt.Println(len(s))
	fmt.Println(cap(s))
	if s == nil {
		fmt.Println("s is nil.")
	} else {
		fmt.Println("s is not nil.")
	}

	//s2 := s[:3]
	//fmt.Println(s2)
}
