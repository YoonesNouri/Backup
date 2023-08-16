// Optimized by removing extra iterations. Mentioned by the word <<here>>.
package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Enter unequal numbers:\n")
	numbers := []int{}
	for {
		fmt.Print("Enter a number: ")
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}
		numbers = append(numbers, n)
	}

	fmt.Println("Sorting from small to big...")
	n = len(numbers)
	for i := 0; i < n-1; i++ { //<<here>>
		swapped := false
		for j := 0; j < n-1-i; j++ { //<<here>>
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		if !swapped {
			break //<<here>> break checks if any swaps were made.
			//If no swaps were made,
			//it means that the slice is already sorted. so breaks out.
		}
	}

	fmt.Println("Sorted:", numbers)
}
