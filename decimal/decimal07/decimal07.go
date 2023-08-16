package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println(BcCmpFloat64(3.14, 3.14))
	fmt.Println(BcCmpFloat64(3.141, 3.143))
	fmt.Println(BcCmpFloat64(3.145, 3.147))
	fmt.Println(BcCmpFloat64(3.141, 3.147))
}

func BcCmpFloat64(num1, num2 float64) int {
	dec1 := decimal.NewFromFloat(num1).Round(2)
	dec2 := decimal.NewFromFloat(num2).Round(2)
	return dec1.Cmp(dec2)
}