package main

import (
	"time"
	"fmt"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

var dilbert Employee

func main() {
	fmt.Println(dilbert)
	dilbert.Salary += 5000
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Printf("%v", dilbert)
}