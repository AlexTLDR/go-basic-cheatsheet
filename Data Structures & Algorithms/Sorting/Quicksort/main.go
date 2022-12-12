package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	list := rand.Perm(50)
	fmt.Println("List before sorting:", list)
	fmt.Println("List after sorting:", quickSort(list))
}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	less, greater := 0, len(list)-1
	//pivot := rand.Int() % len(list)
	pivot := 0
	list[pivot], list[greater] = list[greater], list[pivot]

	for i := range list {
		if list[i] < list[greater] {
			list[i], list[less] = list[less], list[i]
			less++
		}
	}
	list[less], list[greater] = list[greater], list[less]
	quickSort(list[:less])
	quickSort(list[less+1:])

	return list
}
