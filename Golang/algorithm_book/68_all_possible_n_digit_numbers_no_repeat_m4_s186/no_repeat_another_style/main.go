package main

import (
	"fmt"
)

func main() {
	var targetDigit int
	var givenNumbers []int

	// Get target digit from the user
	fmt.Print("Enter the target digit: ")
	fmt.Scanln(&targetDigit)

	// Get given numbers from the user
	fmt.Println("Enter the given numbers (should be more than target digit logically):")
	for {
		var num int
		fmt.Print("Enter a number: ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		givenNumbers = append(givenNumbers, num)
	}

	// Generate and print all possible multi-digit numbers with non-repeating digits
	fmt.Println("Possible multi-digit numbers:")
	count := printNumbers(targetDigit, givenNumbers, []int{})
	fmt.Println("Number of possible multi-digit numbers:", count)
}

func printNumbers(targetDigit int, givenNumbers []int, currentNumber []int) int {
	if targetDigit == 0 {
		fmt.Println(currentNumber)
		return 1
	}

	count := 0
	for _, num := range givenNumbers {
		if contains(currentNumber, num) {
			continue
		}

		currentNumber = append(currentNumber, num)
		count += printNumbers(targetDigit-1, givenNumbers, currentNumber)
		currentNumber = currentNumber[:len(currentNumber)-1]
	}

	return count
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}
