package main

import (
	"github.com/dabao1989/golanglearn/shared-variable/bank"
	"fmt"
)

func main() {
	ch := make(chan struct{}, 2)

	go func() {
		bank.Deposit(300)
		fmt.Println("=", bank.Balance())
		ch <- struct {}{}
	}()

	go func() {
		bank.Deposit(100)
		ch <- struct {}{}
	}()

	for i := 0; i<2; i++ {
		<-ch
	}
}
