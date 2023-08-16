package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {
	// Generate a random number between 1 and 100
	//rand.Seed(time.Now().UnixNano()) نیازی به تغییر سید نامبر نداریم زیرا در ورژن جدید گو خودش تغییر میده
	secretNumber := rand.Intn(100)

	fmt.Println("Welcome to the Guess the Number Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	var guess int
	numGuesses := 0

	for {
		fmt.Print("Take a guess: ")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		numGuesses++

		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the number in %d guesses.\n", numGuesses)
			break
		} else if guess < secretNumber {
			fmt.Println("Too low! Guess higher.")
		} else {
			fmt.Println("Too high! Guess lower.")
		}
	}

	fmt.Println("Thank you for playing the Guess the Number Game!")
}
