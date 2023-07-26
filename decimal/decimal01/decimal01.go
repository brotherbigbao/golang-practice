package main

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func main() {
	price, err := decimal.NewFromString("3.10");
	if err != nil {
		println(err)
	}
	println(price.Float64())

	price2, err := strconv.ParseFloat("3.10", 64)
	if err != nil {
		println(err)
	}
	println(price2)
}
