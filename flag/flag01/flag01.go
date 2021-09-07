package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	//直接返回指针
	wordPtr := flag.String("world", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	//提前定义变量
	var email string
	flag.StringVar(&email, "email", "", "a string")

	//自定义类型 自定义类型要实现Set和String方法
	var pop percentage
	flag.Var(&pop, "pop", "popularity")

	//解析命令行参数
	flag.Parse()

	//打印
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("email:", email)
	fmt.Println("pop:", pop)
	fmt.Println("未定义参数:", flag.Args())
}

type percentage float32

//Set 将命令行的string类型转为指定类型
func (p *percentage) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	*p = percentage(v)
	return err
}
//String 将实际类型转为字符串类型用于展示
func (p *percentage) String() string{
	return fmt.Sprintf("%f", *p)
}
