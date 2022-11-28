package main

import "fmt"

func main() {
	fmt.Println(inRange(1, 100, -1, 34, 24, 121, 89.5))
}

func inRange(min, max float64, numbers ...float64) []float64 {
	result := []float64{}
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}
