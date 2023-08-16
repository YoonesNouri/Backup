package main

import "fmt"

var n int

func perfectnumber(n int) int {
	divisors := []int{}
	sum := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	for _, divisor := range divisors {
		if divisor < n {
			sum += divisor
		}
	}

	if sum == n {
		fmt.Println(n, "is a perfect number.")
	} else {
		fmt.Println(n, "is not a whole number.")
	}
	return sum
}

func main() {
	fmt.Print("enter a number:")
	fmt.Scanln(&n)
	perfectnumber(n)

}
