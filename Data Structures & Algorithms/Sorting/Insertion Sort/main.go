package main

import "fmt"

func Less(a, b int) bool {
	return a < b
}

func Greater(a, b int) bool {
	return a > b
}

func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5, 0}
	InsertionSort(data, Greater)
	fmt.Println(data)
	InsertionSort(data, Less)
	fmt.Println(data)
}
