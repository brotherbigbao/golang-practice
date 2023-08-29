package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	decimalRoundDebug("3.unique_id")
	fmt.Println("==========")
	decimalRoundDebug("3.05")
	fmt.Println("==========")
	decimalRoundDebug("3.09")
	fmt.Println("==========")
}

func decimalRoundDebug(number string)  {
	d, err := decimal.NewFromString(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d.Round(1).String())
	fmt.Println(d.RoundUp(1).String())
	fmt.Println(d.RoundDown(1).String())
	return
}
