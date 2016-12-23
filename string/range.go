package main

import "fmt"

func main() {
	str := "这是一个字符串abc"
	for k, v := range str {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println(string(v))
		fmt.Println("===========")
	}
}
