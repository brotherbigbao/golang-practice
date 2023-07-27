package main

import "github.com/shopspring/decimal"

func main() {
	price, err := decimal.NewFromString("0.0");
	if err != nil {
		println(err)
	}
	println(price.Float64())

	println(price.RoundDown(1).Float64())

	floatNum, _ := price.RoundDown(1).Float64()
	if floatNum == 0 {
		println("=0")
	}
}
