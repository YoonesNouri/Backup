package main

import "fmt"

var quantity, limit, Gnumber int
var numbers []int

func main() {
	fmt.Print("Enter a quantity of numbers from 1 to a limit: \n")

	for {
		fmt.Print("Enter limit: ")
		if _, err := fmt.Scanln(&limit); err != nil {
			fmt.Println("Invalid input. Please enter a valid limit.")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter quantity: ")
		if _, err := fmt.Scanln(&quantity); err != nil {
			fmt.Println("Invalid input. Please enter a valid quantity.")
			continue
		}

		if quantity > 0 && quantity <= limit {
			break
		}

		fmt.Print("The quantity should be within the valid range.\n")
	}

	numbers = make([]int, quantity)

	fmt.Printf("Enter the numbers to be guessed. There are %v numbers within the range of 1 to %v:\n", quantity, limit)
	for i := 0; i < quantity; i++ {
		fmt.Printf("Enter number %v: ", i+1)
		if _, err := fmt.Scanln(&numbers[i]); err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			i-- // Decrement i to allow the user to re-enter the number
			continue
		}
		if numbers[i] > limit {
			fmt.Printf("Number %v exceeds the limit. Please enter a number within the valid range.\n", numbers[i])
			i-- // Decrement i to allow the user to re-enter the number
		}
	}

	fmt.Printf("You have %v chances to guess:\n", quantity*2)
	guesses := 0
	correctGuesses := 0
	wrongGuesses := 0
	allGuessed := false

GameLoop:
	for guesses < quantity*2 {
		if allGuessed {
			break
		}

		fmt.Print("Guess a number (or enter '0' to end the game): ")
		if _, err := fmt.Scanln(&Gnumber); err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		switch Gnumber {
		case 0:
			fmt.Println("Game ended.")
			break GameLoop
		}

		if Gnumber > limit {
			fmt.Printf("The guessed number %v exceeds the limit. Please guess a number within the valid range.\n", Gnumber)
			continue
		}

		found := false
		for i, num := range numbers {
			if Gnumber == num {
				fmt.Println("Bingo!!!!!!!!")
				found = true
				correctGuesses++
				numbers = removeIndex(numbers, i)
				if len(numbers) == 0 {
					allGuessed = true
				}
				break
			}
		}

		if !found {
			fmt.Println("Not this time!")
			wrongGuesses++
		}

		guesses++
	}

	fmt.Printf("Correct guesses: %v\n", correctGuesses)
	fmt.Printf("Wrong guesses: %v\n", wrongGuesses)

	if correctGuesses > quantity/2 {
		fmt.Println("Congratulations! You are a winner!")
	} else {
		fmt.Println("Oops! You didn't guess enough correct numbers. Better luck next time!")
	}
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
