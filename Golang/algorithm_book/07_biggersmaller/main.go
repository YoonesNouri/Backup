package main

import "fmt"

var a, b float64

func bs(a, b float64) {
	if a > b {
		fmt.Printf("%v is bigger than %v \n", a, b)
	} else {
		fmt.Printf("%v is bigger than %v \n", b, a)
	}
}

func main() {
	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&a, &b)

	for a == b {
		fmt.Print("Enter two other numbers: ")
		fmt.Scanln(&a, &b)
	}

	bs(a, b)
}
