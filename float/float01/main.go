package main

import (
	"fmt"
	"math"
)

func main()  {
	a := 1100.1
	b := a * 100
	fmt.Println(b)
	fmt.Println(RoundFloat(b, 2))
}

func RoundFloat(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}
