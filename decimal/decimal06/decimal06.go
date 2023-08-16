package main

import (
	"fmt"
	"math"
)
const MIN = 0.000001
// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}
func main() {
	a := 0.9
	b := 1.0

	if IsEqual(a, b) {
		fmt.Println("a==b")
	}else{
		fmt.Println("a!=b")
	}

	if IsEqualFloat64(a, b) {
		fmt.Println("equal")
	}

	if 0.1 < 0.1 {
		fmt.Println("0.1 < 0.1")
	} else {
		fmt.Println("0.1 = 0.1")
	}
}

func IsEqualFloat64(floatValue1 float64, floatValue2 float64) bool {
	p := 0.00001
	// 判断 floatValue1 与 floatValue2 是否相等
	if math.Dim(floatValue1, floatValue2) < p {
		return true
	} else {
		return false
	}
}