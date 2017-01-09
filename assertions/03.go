package main

import "fmt"

func main() {
	var x interface{}
	x = "This is a string."
	switch x.(type) {
	case nil: fmt.Println("is nil")
	case int,uint: fmt.Println("is int or unint")
	case bool: fmt.Println("is bool")
	case string: fmt.Println("is string")
	default:
		fmt.Printf("don't know")
	}
}


