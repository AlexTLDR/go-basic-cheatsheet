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

// when there is no more swap in one pass of the outer loop, the array is already sorted
func BubbleSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := 1
	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
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
	BubbleSort(data, Greater) // to sort the data ascending
	fmt.Println(data)
	BubbleSort(data, Less) // to sort data descending
	fmt.Println(data)
	BubbleSort2(data, Greater)
	fmt.Println(data)
}
