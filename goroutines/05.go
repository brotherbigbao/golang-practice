package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i<100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for {
			t, ok := <-naturals
			if !ok {
				break
			}
			squares <- t * t
		}
		close(squares)
	}()

	for {
		t, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(t)
	}
}
