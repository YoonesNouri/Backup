package main

import "fmt"

func main() {
	for n := 2; n <= 500; n++ {
		divisors := []int{}
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				divisors = append(divisors, i)
			}
		}
		fmt.Printf("divisors of %d : %v\n", n, divisors)
	}
}
