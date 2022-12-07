package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	//go retrieveWebPage("https://example.com")
	go responseSize("https://example.com")
	go responseSize("https://google.com")
	go responseSize("https://github.com")
	go responseSize("https://gsp.ro")
	time.Sleep(5 * time.Second)
}

func retrieveWebPage(url string) {
	fmt.Println("Displaying", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}
