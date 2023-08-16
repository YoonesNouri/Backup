package main

import "fmt"

var a, b int

func BMM(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	fmt.Print("enter first number:")
	fmt.Scanln(&a)
	fmt.Print("enter second number:")
	fmt.Scanln(&b)

	bmm := BMM(a, b)
	fmt.Println("BMM of", a, "and", b, "is", bmm)
	kmm := a * b / bmm
	fmt.Println("KMM of", a, "and", b, "is", kmm)
}
