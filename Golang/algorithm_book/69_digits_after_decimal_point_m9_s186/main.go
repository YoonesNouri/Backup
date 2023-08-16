package main

import (
	"fmt"
	"strings"
)

func main() {
	var dividend, divisor, desiredDigit int

	// Get dividend from the user
	fmt.Print("Enter the dividend: ")
	fmt.Scanln(&dividend)

	// Get divisor from the user
	fmt.Print("Enter the divisor: ")
	fmt.Scanln(&divisor)

	// Get desired digit from the user
	fmt.Print("Enter the desired digit: ")
	fmt.Scanln(&desiredDigit)

	// Perform division and calculate digits
	digits := calculateDigits(dividend, divisor, desiredDigit)

	// Print the digits after the decimal point
	fmt.Printf("Digits after decimal point up to digit %d: %s\n", desiredDigit, digits)
}

func calculateDigits(dividend, divisor, desiredDigit int) string {
	// Perform long division
	quotient := dividend / divisor
	remainder := dividend % divisor
	decimalDigits := []string{}

	// Iterate to calculate desired digits
	for i := 0; i < desiredDigit; i++ {
		remainder *= 10
		quotient = remainder / divisor
		remainder = remainder % divisor
		decimalDigits = append(decimalDigits, fmt.Sprintf("%d", quotient))
	}

	decimalDigitsStr := strings.Join(decimalDigits, "")

	return decimalDigitsStr
}
