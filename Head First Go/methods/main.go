package main

import "fmt"

type liters float64
type milliliters float64
type gallons float64

func main() {
	soda := liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.toGallons())
	water := milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.toGallons())
}

func (l liters) toGallons() gallons {
	return gallons(l * 0.264)
}

func (m milliliters) toGallons() gallons {
	return gallons(m * 0.000264)
}
