package main

import (
	"fmt"
	"time"
)

func main() {
	go first()
	go two()
	time.Sleep(200 * time.Millisecond)
}

func first() {
	for i := 0; i<10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond)
	}
}

func two() {
	for i := 0; i<10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond)
	}
}
