package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["liuyibao"] = 28
	ages["huiting"] = 26
	ages["ze"] = 23

	//var names []string
	for name := range ages {
		fmt.Println(name)
	}
	for name, age := range ages {
		fmt.Println(name, age)
	}

	s := [...]string{"abc", "def", "hello"}
	for key := range s {
		fmt.Println(key)
	}
	for key, str := range s {
		fmt.Println(key, str)
	}
}
