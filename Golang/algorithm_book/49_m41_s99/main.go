package main

import (
	"fmt"
)

func main() {
	var a, b int

	for {
		fmt.Print("enter a: ")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Print("invalid input.")
			continue
		}
		break
	}
	for {
		fmt.Print("enter b: ")
		_, err := fmt.Scanln(&b)
		if err != nil {
			fmt.Print("invalid input.")
			continue
		}
		break
	}

	multiply := 0
	for i :=1 ; i <=a ; i++{
multiply += b
	}

	fmt.Printf("a x b = %v", multiply)
}
