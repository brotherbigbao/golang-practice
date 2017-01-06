package main

import (
	"fmt"
	"image/color"
)

func main() {
	p := Point{2, 4}
	fmt.Println(p)

	red := color.RGBA{255, 0, 0, 255}
	cp := ColoredPoint{&p, red}
	fmt.Println(cp)
	fmt.Println(cp.Point)
}