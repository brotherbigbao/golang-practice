package main

import "fmt"

func main() {
	if 1==1 {
		defer func() {
			fmt.Println("defer 已执行")
		}()
	}
	fmt.Println("测试 defer")
}
