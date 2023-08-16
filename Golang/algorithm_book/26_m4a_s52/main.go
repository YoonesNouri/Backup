package main

import (
	"fmt"
)

var num int
var sum float64

func main() {
	fmt.Print("enter a number:")
	fmt.Scanln(&num)

	sum = 0.0
	for i := 1; i <= num; i++ {
		if i%3 == 0 {
			sum += 1.0 / float64(i)
		}
	}
	fmt.Println("final sum = ", sum)
}
