package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"fmt"
)

var cacheCinemas []byte

func main () {
	go cleanCache()
	http.HandleFunc("/movie/v3/cinemas", cinemas)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func cinemas (w http.ResponseWriter, r *http.Request) {
	if len(cacheCinemas) > 0 {
		w.WriteHeader(200)
		w.Write(cacheCinemas)
	} else {
		resp, err := http.Get("https://api.ffan.com/movie/v3/cinemas?cityId=110100")
		if err != nil {
			log.Fatal("内部代理错误")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		cacheCinemas = body;

		w.WriteHeader(200)
		w.Write(body)
	}
}

func cleanCache () {
	for {
		time.Sleep(10000 * time.Millisecond)
		fmt.Println(cacheCinemas)
		cacheCinemas = []byte{}
		fmt.Println(cacheCinemas)
	}
}