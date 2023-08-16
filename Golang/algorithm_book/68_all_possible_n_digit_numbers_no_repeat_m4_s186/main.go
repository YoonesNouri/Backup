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

	// Generate all possible multi-digit numbers without repetition
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
	// Recursive function to generate multi-digit numbers without repetition
	var generateHelper func(digit int, currentNumber []int, used []bool, result *[][]int)

	generateHelper = func(digit int, currentNumber []int, used []bool, result *[][]int) {
		if digit == 0 {
			*result = append(*result, append([]int(nil), currentNumber...))
			return
		}

		for i, num := range givenNumbers {
			if !used[i] {
				used[i] = true
				generateHelper(digit-1, append(currentNumber, num), used, result)
				used[i] = false
			}
		}
	}

	// Generate multi-digit numbers without repetition
	numbers := [][]int{}
	used := make([]bool, len(givenNumbers))
	generateHelper(targetDigit, []int{}, used, &numbers)

	return numbers
}
