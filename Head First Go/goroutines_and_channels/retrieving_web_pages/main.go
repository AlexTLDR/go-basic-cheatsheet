package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//go retrieveWebPage("https://example.com")

	sizes := make(chan int)

	go responseSize("https://example.com", sizes)
	go responseSize("https://google.com", sizes)
	go responseSize("https://github.com", sizes)
	go responseSize("https://gsp.ro", sizes)

	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
}

// func retrieveWebPage(url string) {
// 	fmt.Println("Displaying", url)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(body))
// }

func responseSize(url string, channel chan int) {
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
	channel <- len(body)
}
