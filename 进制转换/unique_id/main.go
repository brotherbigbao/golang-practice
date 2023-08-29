package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number uint64 = 2292550984142401
	str := strconv.FormatUint(number, 2)
	//原始二进制
	fmt.Println(str)

	prefix := ""
	if 64 - len(str) > 0 {
		fillZeroNum := 64 - len(str)
		for i:=0; i<fillZeroNum; i++ {
			prefix += "0"
		}
	}
	str = prefix + str
	//补全64位
	fmt.Println(str)

	fmt.Println(str[0:11], " 11位空缺")
	fmt.Println(str[0:40], " 11位空缺+29位时间戳")
	fmt.Println(str[0:45], " 11位空缺+29位时间戳+5位业务线")
	fmt.Println(str[0:48], " 11位空缺+29位时间戳+5位业务线+3位机器ID")
	fmt.Println(str[0:64], " 11位空缺+29位时间戳+5位业务线+3位机器ID+16位随机数")
}
