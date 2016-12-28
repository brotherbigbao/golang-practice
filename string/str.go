package main

import "fmt"

func main(){
	s := "你好，世界"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println(s[20])
	fmt.Println("hello")
}
