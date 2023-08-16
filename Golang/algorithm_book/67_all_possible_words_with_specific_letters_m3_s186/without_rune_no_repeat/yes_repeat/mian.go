package main

import "fmt"

func main() {
	letters := []string{"I", "R", "A", "N"}
	words := generateWords(letters)

	fmt.Println("Possible words:")
	for _, word := range words {
		fmt.Println(word)
	}

	fmt.Println("Number of possible words:", len(words))
}

func generateWords(letters []string) []string {
	words := []string{}
	word := ""
	backtrack(letters, word, &words)

	return words
}

func backtrack(letters []string, word string, words *[]string) {
	if len(word) == len(letters) {
		*words = append(*words, word)
		return
	}

	for i := 0; i < len(letters); i++ {
		word += letters[i]
		backtrack(letters, word, words)
		word = word[:len(word)-1]
	}
}
