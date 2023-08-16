package main

import (
	"fmt"
)

var n, numdigit, sumdigit int
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
	fmt.Print("Enter number: ")
	fmt.Scanln(&n)

	digits = splitInt(n)
	numdigit = len(digits)

	sumdigit = 0
	for _, v := range digits {
		sumdigit += v
	}

	evendigits := []int{}
	for _, v := range digits {
		if v%2 == 0 {
			evendigits = append(evendigits, v)
		}
	}
	numevendigits := len(evendigits)

	odddigits := []int{}
	for _, v := range digits {
		if v%2 != 0 {
			odddigits = append(odddigits, v)
		}
	}
	numodddigits := len(odddigits)

	zerodigits := []int{}
	for _, v := range digits {
		if v == 0 {
			zerodigits = append(zerodigits, v)
		}
	}
	numzerodigits := len(zerodigits)

	fmt.Printf("Digits of %v: %v\n", n, digits)
	fmt.Printf("Number of digits: %d\n", numdigit)
	fmt.Printf("Sum of digits: %d\n", sumdigit)

	fmt.Printf("evendigits: %d\n", evendigits)
	fmt.Printf("number of evendigits: %d\n", numevendigits)

	fmt.Printf("odddigits: %d\n", odddigits)
	fmt.Printf("number of odddigits: %d\n", numodddigits)

	fmt.Printf("number of zerodigits: %d\n", numzerodigits)

}
