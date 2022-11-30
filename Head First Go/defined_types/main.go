package main

import "fmt"

type liters float64
type gallons float64

func main() {
	fmt.Println(toGallons(10))
	fmt.Println(toLiters(10))
}

func toGallons(l liters) gallons {
	return gallons(l * 0.264)
}

func toLiters(g gallons) liters {
	return liters(g * 3.785)
}
