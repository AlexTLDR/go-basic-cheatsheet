package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum(34.5, 101.5, 23, 98.7))
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}

	}
	return max
}
