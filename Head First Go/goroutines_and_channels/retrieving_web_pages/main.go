package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	retrieveWebPage("https://example.com")
}

func retrieveWebPage(url string) {
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
