package main

import (
	"log"
	"net/http"
)

func httpServer() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe("", nil) // listen to "http://localhost:80"
	if err != nil {
		log.Fatal("Http listen and server error", err)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("In - helloHandler - ", req.Method, " URL:", req.URL)
	w.Write([]byte("hello world"))
	log.Println("Out - helloHandler")
}

/*
When hit get request in Brower (or) run the httpClient.go file from other main project
http://localhost:80/hello
Output:
hello world
*/
