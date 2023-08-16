package main

import (
	"fmt"
)

var n, sum int

func main() {

	fmt.Print("Enter limit (n>0):")
	for {
		fmt.Scanln(&n)
		if n > 0 {
			break
		}
		fmt.Print("Enter limit (n>0):")
	}

	//Sum = 1*2 + 2*3 + ... + n*(n+1)
	sum = 0
	for i := 1; i <= n; i++ {
		sum += i * (i + 1)
		fmt.Printf("sum till n(%v) = %v \n", i, sum)
	}
}
