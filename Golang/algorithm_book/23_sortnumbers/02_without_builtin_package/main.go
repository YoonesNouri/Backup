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
	for i := 0; i < n; i++ { //دو لوپ نیاز است تا مطمئن شویم همه ی عناصر واجد شرط، جابجا شده اند
		for j := 0; j < n-1; j++ { //چرا تا اِن-1؟ چون مقایسه بین اِن عنصر از یک اسلایس، اِن-1 بار است
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Println("Sorted:", numbers)
}
