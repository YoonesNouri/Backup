package main

import (
	"fmt"
	"github.com/YoonesNouri/mymodule/digits"
)

var num int

func main() {
	for {
		fmt.Print("enter a number: ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Printf("invalid input")
			continue
		}
		break
	}

	reverseNumber := digits.RevSplitInt(num)
	fmt.Println("reverseNumber =", reverseNumber)
}
