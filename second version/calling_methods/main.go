package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	year := now.Year()
	fmt.Println(year)
}
