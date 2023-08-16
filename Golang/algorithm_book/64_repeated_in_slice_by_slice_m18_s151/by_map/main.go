package main

import (
	"fmt"
)

func main() {
	numbers := []int{}
	counts := map[int]int{}

	for {
		var num int
		fmt.Print("Enter a number (or enter any non-number value to stop): ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
		counts[num]++
	}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Repeated Numbers:")
	for num, count := range counts {
		if count > 1 {
			fmt.Printf("%v is repeated %v times\n", num, count)
		}
	}
}
