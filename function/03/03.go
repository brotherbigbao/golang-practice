package main

import "fmt"

func main() {
	a := 1;
	b := 2;
	c := 3;
	fmt.Println(toSlice(&a, &b, &c))
}

func toSlice(i ...interface{}) []interface{} {
	return i
}