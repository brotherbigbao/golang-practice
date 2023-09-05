package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

var data = `
redis:
  addr: "127.0.0.1:6379"
  password: "123456"
  db: 0
  MaxIdle: 10
  MaxActive: 10
  IdleTimeout: 10
  Wait: true
  MaxConnLifetime: 600

db:
  name: "prod"
  maxConnNum: 15
  maxIdleConns: 5
  master:
    addr: "127.0.0.1:3306"
    user: "root"
    password: "123456"
    weight: 1
  slave:
    -
      addr: "127.0.0.1:3306"
      user: "root"
      password: "123456"
      weight: 1
    -
      addr: "127.0.0.1:3306"
      user: "root"
      password: "123457"
      weight: 1
    -
      addr: "127.0.0.1:3306"
      user: "root"
      password: "123458"
      weight: 1
`

func main() {
	var t map[string]interface{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t["db"])
	/**
	slave:[
	map[addr:127.0.0.1:3306 password:123456 user:root weight:1]
	map[addr:127.0.0.1:3306 password:123457 user:root weight:1]
	map[addr:127.0.0.1:3306 password:123458 user:root weight:1]
	]
	*/

}
