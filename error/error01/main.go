package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := errors.New("This is a error")
	fmt.Println(err)

	newErr := fmt.Errorf("wrap了一个错误%w", err)
	fmt.Println(newErr)
	errors.Is(newErr, os.ErrNotExist)
}
