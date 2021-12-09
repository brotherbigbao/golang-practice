package main

import "fmt"

func main(){
	a := sum(7,8,9,10)
	b := sum(1,2,3,4,5)
	fmt.Println(a)
	fmt.Println(b)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
