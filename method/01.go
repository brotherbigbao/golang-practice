package main

import (
	"github.com/dabao1989/golanglearn/method/geometry"
	"fmt"
)

func main() {
	a := geometry.Point{3.33334, 6.8432423}
	b := geometry.Point{12.342342134, 9.3423424}
	distance := a.Distance(b)
	fmt.Println(distance)
}
