package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	data ,err := ioutil.ReadFile("./tmp.json")
	if err != nil {
		fmt.Println(data)
		return
	}

	baser64Body := base64.URLEncoding.EncodeToString(data)
	//baser64Body := base64.StdEncoding.EncodeToString(data)

	fmt.Println(baser64Body)
}
