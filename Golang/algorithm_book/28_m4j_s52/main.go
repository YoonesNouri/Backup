package main

import (
	"fmt"
	"math"
)

var x, n int
var sum float64

func fac(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Print("enter a number (x>0):")
	for {
		fmt.Scanln(&x)
		if x > 0 {
			break
		}
		fmt.Print("enter a number (x>0):")
	}
	fmt.Print("enter limit power (n>0):")
	for {
		fmt.Scanln(&n)
		if n > 0 {
			break
		}
		fmt.Print("enter limit power (n>0):")
	}

	//Sum = x + x^3 / 3! + ... + x^(2n-1) / (2n-1)!
	sum = 0.0
	for i := 1; i <= n; i++ {
		sum += math.Pow(float64(x), float64(2*i-1)) / float64(fac(2*i-1))
		fmt.Printf("sum of %v powered till %v = %v \n", x, i, sum)
	}
}
