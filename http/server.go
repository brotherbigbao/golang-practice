package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/name", name)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello"))
	fmt.Fprintf(w, "Hello astaxie!")
}

func name(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("this is name"))
}