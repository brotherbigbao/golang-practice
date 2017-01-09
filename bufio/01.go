package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println(input.Text())

}
