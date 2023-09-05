package main

import "fmt"

func main() {
	var a []*int

	v1 := 1
	v2 := 2
	v3 := 3
	v4 := 4
	v5 := 5

	a = append(a, &v1, &v2, &v3, &v4, &v5)

	for _, v := range a {
		fmt.Println(*v)
	}

	for _, v := range a {
		*v = *v * 3
	}
}
