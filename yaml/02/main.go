package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

var data = `
slave:
- addr: "127.0.0.1:3306"
  user: "root"
  password: "123456"
  weight: 1
- addr: "127.0.0.1:3306"
  user: "root"
  password: "123456"
  weight: 1
- addr: "127.0.0.1:3306"
  user: "root"
  password: "123456"
  weight: 1
`

func main() {
	var t map[string]interface{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t["slave"])
	/**
	[
	map[addr:127.0.0.1:3306 password:123456 user:root weight:1]
	map[addr:127.0.0.1:3306 password:123456 user:root weight:1]
	map[addr:127.0.0.1:3306 password:123456 user:root weight:1]
	]
	 */
}
