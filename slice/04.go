package main

import "fmt"

func main() {
	s := make([]string, 10, 1000)
	for k,v := range s {
		fmt.Println(k, v)
	}

	//s = s[:1001]
	//fmt.Println(len(s))

	s = s[:1000]
	fmt.Println(len(s))
}
