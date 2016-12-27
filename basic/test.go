package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	a["liuyibao"] = 3

	age, ok := a["liuyibao"]
	if ok {
		fmt.Println(age)
	}

	age, ok = a["other"]
	if ok {
		fmt.Println(age)
	}
}