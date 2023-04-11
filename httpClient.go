package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// use another main project and run this httpClient to get the response
func httpClient() {
	resp, err := http.Get("http://localhost:80/hello")
	if err != nil {
		log.Fatal("Error received for http get call")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error received for io read")
	}
	fmt.Println(string(body))
}
