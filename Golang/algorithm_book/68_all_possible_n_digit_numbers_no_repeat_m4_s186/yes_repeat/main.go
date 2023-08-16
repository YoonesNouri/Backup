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

	// Generate all possible multi-digit numbers with repetition
	possiblenumbers := generateNumbers(targetDigit, givenNumbers)

	// Print the generated multi-digit numbers
	fmt.Println("Possible multi-digit numbers:")
	for _, row := range possiblenumbers {
		fmt.Println(row)
	}

	// Print the number of generated multi-digit numbers
	fmt.Println("Number of possible multi-digit numbers:", len(possiblenumbers))
}

func generateNumbers(targetDigit int, givenNumbers []int) [][]int {
	// Generate multi-digit numbers with repetition
	numbers := [][]int{}
	currentNumber := make([]int, targetDigit)

	var backtrack func(digit int)
	backtrack = func(digit int) {
		if digit == 0 {
			// Append a copy of the current number to the result
			temp := make([]int, targetDigit)
			copy(temp, currentNumber)
			numbers = append(numbers, temp)
			return
		}

		for _, num := range givenNumbers {
			currentNumber[targetDigit-digit] = num
			backtrack(digit - 1)
		}
	}

	backtrack(targetDigit)

	return numbers
}
