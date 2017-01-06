package main

import "fmt"

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Catch a panic: %v", p)
		} else {
			fmt.Println("Correct.")
		}
	}()
	for i := 3; i>=0; i-- {
		fmt.Printf("8/i=%v\n", 8.0/i)
	}
}