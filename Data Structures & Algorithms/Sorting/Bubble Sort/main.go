package main

import "fmt"

func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func Less(a, b int) bool {
	return a < b
}

func Greater(a, b int) bool {
	return a > b
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data, Greater)
	fmt.Println(data)
}
