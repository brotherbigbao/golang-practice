package main

import (
	"github.com/dabao1989/golanglearn/method/animal"
	"fmt"
)

func main() {
	cat := animal.Cat{}
	cat.SetAge(5)
	cat.SetColor("gray")
	cat.SetName("maomi")
	cat.SetPrice(23.345)
	fmt.Println(cat)
}
