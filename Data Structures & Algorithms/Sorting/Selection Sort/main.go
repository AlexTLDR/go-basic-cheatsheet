package main

import "fmt"

// putting the largest value at the end
func SelectionSort(arr []int) {
	size := len(arr)
	var i, j, max int
	for i = 0; i < size; i++ {
		max = 0
		for j = 1; j < size-1-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
	}
}

// putting the smallest value at the begining
func SelectionSort2(arr []int) {
	size := len(arr)
	var i, j, min int
	for i = 0; i < size-1; i++ {
		min = i
		for j = i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	data2 := []int{98, 35, 12, 87, 100, 57, 13, 65, 32}
	SelectionSort(data)
	fmt.Println(data)
	SelectionSort2(data2)
	fmt.Println(data2)

}
