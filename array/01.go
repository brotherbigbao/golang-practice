package main

import "fmt"

func main() {
	var q [3]int = [...]int{1,2,3}
	fmt.Println(q)
	fmt.Printf("%T\n", q)

	symbol := [...]string{0: "name", 1:"sex", 3:"age"}
	fmt.Println(symbol[2])
	fmt.Println(len(symbol))

	fmt.Println("=================")
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [3]int{1,2}
	fmt.Println(a == d)
}
