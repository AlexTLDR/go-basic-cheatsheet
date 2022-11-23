package main

import (
	"fmt"
)

func main() {

	// a slice expression doesn't create a new backing array. The original and the returned slice are connected!
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] //s3, s4 share the same backing array with s1

	s3[1] = 600     // modifying the backing array so s1, s3 and s4 are in fact modified!!
	fmt.Println(s1) // -> [10 600 30 40 50]
	fmt.Println(s4) // -> [600 30]

	// when a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2                       // modifying the array
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Subaru", "Isuzu", "Mazda", "Toyota"}
	newCars := []string{}

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                              // only cars is modified
	fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Isuzu Mazda Toyota] newCars: [Subaru Isuzu

}
