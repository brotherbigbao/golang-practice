package main

import (
	"os"
	"strings"
	"fmt"
	"bufio"
	"strconv"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("请指定文件名和行号, eg: comment index.php 7,8,9")
		os.Exit(0)
	}

	lines := strings.Split(os.Args[2], ",")
	var lineNums []int
	for _, line := range lines {
		lineNum, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(101)
		}
		lineNum--
		lineNums = append(lineNums, lineNum)
	}

	var content []string
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(100)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		content = append(content, input.Text())
	}

	fileLineNums := len(content) - 1
	for _, lineNum := range lineNums {
		if lineNum > fileLineNums {
			lineNum++
			fmt.Fprintf(os.Stderr, "lineNum %d is too max.", lineNum)
			os.Exit(102)
		}

		if lineNum < 0 {
			lineNum++
			fmt.Fprintf(os.Stderr, "lineNum %d is too min.", lineNum)
			os.Exit(103)
		}

		if len(os.Args) == 4 && os.Args[3] == "-r" {
			content[lineNum] = strings.TrimPrefix(content[lineNum], "//")
		}else if !strings.HasPrefix(content[lineNum], "//"){

			content[lineNum] = "//" + content[lineNum]
		}

	}

	newFileContent := strings.Join(content, "\n")
	ioutil.WriteFile(os.Args[1], []byte(newFileContent), 0664)
}