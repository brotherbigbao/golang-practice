package main

import (
	"fmt"
	"sort"
)

func main() {
	var lockSns []int
	lockSns = append(lockSns, 8, 9, 133, 0, 5, 7)
	sort.Sort(sort.IntSlice(lockSns))

	fmt.Println(lockSns)
}
