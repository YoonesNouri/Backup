package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generateWords(letters []rune, word []rune, used []bool, index int, count *int) {
	if index == len(word) {
		fmt.Println(string(word))
		*count++
		return
	}

	for i := 0; i < len(letters); i++ {
		if !used[i] {
			word[index] = letters[i]
			used[i] = true
			generateWords(letters, word, used, index+1, count)
			used[i] = false
		}
	}
}

func main() {
	letters := []rune{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter letters to be used in possible words (enter blank line to stop):")
	for {
		fmt.Print("Enter a letter: ")
		scanner.Scan()
		lett := strings.TrimSpace(scanner.Text())
		if lett == "" {
			break
		}
		runeLetters := []rune(lett)
		if len(runeLetters) != 1 {
			fmt.Println("Invalid input. Please enter a single letter.")
			continue
		}
		letters = append(letters, runeLetters[0])
	}
	wordLength := len(letters)
	word := make([]rune, wordLength)
	used := make([]bool, len(letters))
	count := 0

	generateWords(letters, word, used, 0, &count)

	fmt.Println("Number of possible words:", count)
}
