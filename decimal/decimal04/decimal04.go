package main

import "github.com/shopspring/decimal"

func main() {
	val := BcSubFloat64(3.2, 3.1, 0)
	if val == 0 {
		println("equal 0")
	} else {
		println("not equal 0")
	}
}

//浮点数float64 安全减法 注意会四舍五入 如果不是最终计算结果请多保留几位小数
func BcSubFloat64(num1, num2 float64, precision int32) float64 {
	dec1 := decimal.NewFromFloat(num1)
	dec2 := decimal.NewFromFloat(num2)
	floatNum, _ := dec1.Sub(dec2).Round(precision).Float64()
	return floatNum
}