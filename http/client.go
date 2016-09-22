package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal("请求错误")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取body错误")
	}
	for _, v := range body {
		fmt.Print(string(v))
	}
}
