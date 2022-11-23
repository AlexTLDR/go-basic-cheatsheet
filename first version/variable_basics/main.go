package main

import "fmt"

func main() {
	var age = 34
	fmt.Println("Age: ", age)

	var name = "Dan"
	//fmt.Println("Your name is:", name)
	_ = name

	s := "Gopher"
	fmt.Println(s)

	name = "Alex"
	name1 := "John"
	_ = name1

	car, cost := "Isuzu", 5000
	fmt.Println(car, cost)
	car, year := "Mazda", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int

	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	j, i = i, j

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

}
