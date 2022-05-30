package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	absPath, err := filepath.Abs("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(absPath)

	fmt.Println(filepath.Dir(absPath))
}
