package main

import "github.com/shopspring/decimal"

func main() {
	price, err := decimal.NewFromString("3.1015926");
	if err != nil {
		println(err)
	}
	println(price.Float64())

	println(price.RoundDown(1).Float64())
}
