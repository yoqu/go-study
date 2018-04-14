package main

import (
	"net/http"
	"log"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/count", counter)
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		lissajous(writer)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}