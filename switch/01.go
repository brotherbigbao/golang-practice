package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1: fmt.Println("is 1")
	case 2: fmt.Println("is 2")
	case 3: fmt.Println("is 3")
	case 4: fmt.Println("is 4")
		default: fmt.Println("this is default")
	}
}