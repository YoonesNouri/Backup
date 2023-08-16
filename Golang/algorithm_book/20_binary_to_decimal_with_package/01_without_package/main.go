package main

import (
	"fmt"
)

var digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func convertBase(num, fromBase, toBase int) string {
	if num == 0 {
		return "0"
	}

	var result []byte

	n := num
	if n < 0 {
		n = -n
		result = append(result, '-')
	}

	var stack []int
	for n > 0 {
		remainder := n % toBase
		stack = append(stack, remainder)
		n /= toBase
	}

	for len(stack) > 0 {
		result = append(result, digits[stack[len(stack)-1]])
		stack = stack[:len(stack)-1]
	}

	return string(result)
}

func main() {
	var num, fromBase, toBase int

	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)

	fmt.Print("Enter the base of origin: ")
	fmt.Scanln(&fromBase)

	fmt.Print("Enter the base destination: ")
	fmt.Scanln(&toBase)

	// Convert from the original base to decimal first
	decimalNum := 0
	n := num
	p := 1
	for n > 0 {
		remainder := n % 10
		decimalNum += remainder * p
		n /= 10
		p *= fromBase
	}

	// Convert the decimal number to the desired base
	result := convertBase(decimalNum, 10, toBase)

	fmt.Printf("%v based on %v = %v based on %v.\n", num, fromBase, result, toBase)
}
