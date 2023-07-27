package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	str := SprintFloat64(3.01, 1)
	fmt.Println(str)
	if str == "3" {
		println("Yes")
	} else {
		println("no")
	}

	a := fmt.Sprintf("%.1f", 3.16)
	println(a)
}

func SprintFloat64(f float64, precision int32) string {
	d := decimal.NewFromFloat(f)
	return d.RoundDown(precision).String()
}