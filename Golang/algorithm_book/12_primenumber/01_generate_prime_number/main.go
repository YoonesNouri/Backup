package main

import (
	"fmt"
	"math"
)

func isPrimeNumber(num int) bool {
	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func generatePrimeNumbers(limit int) []int {
	primeNumbers := []int{}

	for i := 2; i <= limit; i++ {
		if isPrimeNumber(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}

func main() {
	var limit int
	fmt.Print("Enter limit: ")
	fmt.Scanln(&limit)

	primeNumbers := generatePrimeNumbers(limit)

	fmt.Println("Prime numbers up to", limit, ":")
	for _, num := range primeNumbers {
		fmt.Println(num)
	}
	numprime := len(primeNumbers)
	fmt.Printf("there are %v prime numbers between 2 to %v", numprime, limit)
}
