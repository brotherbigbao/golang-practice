package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["liuyibao"] = 28
	ages["else"] = 27
	fmt.Println(ages)

	heights := map[string]int{}
	heights["liuyibao"] = 174
	heights["else"] = 175
	fmt.Println(heights)

	amount := map[string]int{
		"macbook": 10000,
		"thinkpad": 8000,
	}
	fmt.Println(amount)
	delete(amount, "macbook")
	delete(amount, "something")//即使此键不存在也没有关系
	fmt.Println(amount)
}
