package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Amir"

	result := repeatLetters(word)
	fmt.Println(result) // Output: heelllllllooooo
}

func repeatLetters(word string) string {
	var result strings.Builder
	count := 1

	for _, char := range word {
		for j := 0; j < count; j++ {
			result.WriteRune(char)
		}

		count++
	}

	return result.String()
}
