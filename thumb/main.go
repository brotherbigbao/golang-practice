package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"gopl.io/ch8/thumbnail"
	"strings"
	"time"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _,file := range files {
		//fmt.Println(file.Name())
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".jpg") {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	if len(fileNames) < 1 {
		log.Fatal("No image in this dir.")
	}

	start := time.Now()
	//makeThumb(fileNames)
	makeThumb2(fileNames)
	t := time.Now().Sub(start)
	fmt.Println(t.Seconds())
}

func makeThumb(fileNames []string) {
	for _, f := range fileNames {
		thumbnail.ImageFile(f)
	}
}

func makeThumb2(fileNames []string) {
	ch := make(chan struct{}, len(fileNames))
	for _, f := range fileNames {
		go func (f string) {
			thumbnail.ImageFile(f)
			ch <- struct {}{}
		}(f)
	}
	for range fileNames {
		<-ch
	}
}