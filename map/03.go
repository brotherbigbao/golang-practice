package main

import "fmt"

func main() {
	//map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个 panic异常
	var ages map[string]int
	ages["tmp"] = 3 //error
	fmt.Println(ages)
}
