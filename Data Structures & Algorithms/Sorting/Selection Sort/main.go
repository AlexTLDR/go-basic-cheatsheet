package main

import "fmt"

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

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	SelectionSort(data)
	fmt.Println(data)
}
