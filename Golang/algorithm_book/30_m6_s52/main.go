package main

import (
	"fmt"
)

var n int
var digits []int

func splitInt(n int) []int {
	digits := []int{}
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n = n / 10
	}
	return digits
}

func main() {

	fmt.Print("Enter number (n>0):")
	for {
		fmt.Scanln(&n)
		if n > 0 {
			break
		}
		fmt.Print("Enter number (n>0):")
	}

	digits = splitInt(n)
	sumdigit := 0
	for _, v := range digits {
		sumdigit += v
	}

	if n%sumdigit == 0 {
		fmt.Printf("%v is divisible by the sum of its digits", n)
	} else {
		fmt.Printf("%v is NOT divisible by the sum of its digits", n)
	}
}
