package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-c:
			fmt.Println("closed")
			return
		}
	}

}

func main() {
	c := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, c)
	}

	time.Sleep(time.Second)
	close(c)

	wg.Wait()
}
