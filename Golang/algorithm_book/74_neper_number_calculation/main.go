//عدد نِپِر یا اویلر را محاسبه کنید
//e = 1 + 1/1! + 1/2! + 1/3! + 1/4! + ... + (1/n!)
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	numTerms := 20 // Number of terms to include in the series

	sum := 1.0 // Initialize the sum with the first term (1/0!)
	for i := 1; i <= numTerms; i++ {
		term := 1.0 / float64(factorial(i))
		sum += term
	}

	fmt.Println("Approximation of e:", sum)
}
