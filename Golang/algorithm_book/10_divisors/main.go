package main

import "fmt"

var n int

func getdivisors(n int) ([]int, int, int) {
	divisors := []int{}
	sum := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	numdivisors := len(divisors)

	for _, divisor := range divisors {
		sum += divisor
	}
	return divisors, numdivisors, sum
}

func main() {
	fmt.Print("enter a number:")
	fmt.Scanln(&n)
	divisors, numdivisors, sum := getdivisors(n) //must be just 3 variables and can be other names on the left of the :=
	fmt.Printf("divisors of %d : %v\n", n, divisors)
	fmt.Printf("number of divisors : %v\n", numdivisors)
	fmt.Printf("sum of divisors : %v\n", sum)
}
