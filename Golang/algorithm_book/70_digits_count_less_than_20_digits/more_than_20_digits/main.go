package main

import (
	"fmt"
	"math/big"
)

func main() {
	var numStr string
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&numStr)

	num := new(big.Int)
	_, success := num.SetString(numStr, 10)
	if !success {
		fmt.Println("Invalid integer")
		return
	}

	count := 0
	ten := big.NewInt(10)
	zero := big.NewInt(0)
	for num.Cmp(zero) != 0 {
		num.Div(num, ten)
		count++
	}

	fmt.Println("Number of digits:", count)
}
