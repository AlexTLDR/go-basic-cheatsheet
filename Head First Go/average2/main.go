package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	sum := 0.0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Printf("Average: %0.2f\n", sum/float64(len(arguments)))
}
