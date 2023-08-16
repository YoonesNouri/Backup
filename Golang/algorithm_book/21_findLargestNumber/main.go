package main

import "fmt"

func findLargestNumber(numbers ...int) int {
	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func main() {
	var n int
	fmt.Print("Enter the number of values: ")
	fmt.Scanln(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&numbers[i])
	}

	largest := findLargestNumber(numbers...)
	fmt.Printf("The largest number is: %d\n", largest)
}
