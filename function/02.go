package main

import "fmt"

func main() {
	helloDefer()
}

func helloDefer() {
	fmt.Println("Start helloDefer.")
	defer fmt.Println("This is defer.")
	fmt.Println("End helloDefer.")
}
