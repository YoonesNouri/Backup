package main

import (
	"fmt"
	"math"
)

var x, n, sum int

func main() {
	fmt.Print("enter a number (x>0):")
	for {
		fmt.Scanln(&x)
		if x > 0 {
			break
		}
		fmt.Print("enter a number (x>0):")
	}

	fmt.Print("enter limit power (n>=0):")
	for {
		fmt.Scanln(&n)
		if n >= 0 {
			break
		}
		fmt.Print("enter final power (n>0):")
	}
	//Sum = x^0 + x^1 + x^2 + ... + x^n
	sum = 0
	for i := 0; i <= n; i++ {
		sum += int(math.Pow(float64(x), float64(i)))
		fmt.Printf("sum of %v powered till %v = %v \n", x, i, sum)
	}
}
