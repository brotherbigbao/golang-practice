package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["liuyibao"] = 28
	ages["other"] = 30
	if age, ok := ages["others"]; ok {
		fmt.Println(age)
	} else {
		fmt.Println("not exist.")
	}
}
