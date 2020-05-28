package main

import (
	"fmt"
	"time"
)

func work(s chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
		case <-s:
			fmt.Println("closed")
			return
		}
	}
}

func main()  {
	s := make(chan bool)
	go work(s)

	time.Sleep(time.Second)
	//close(s)
	s <- true
}
