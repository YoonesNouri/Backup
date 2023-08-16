package main

import "fmt"

var a int

func main() {
	fmt.Print("enter a number:")
	fmt.Scanln(&a)
	if a == 0 || a%2 == 0 {
		fmt.Printf("%v is even", a)
	} else {
		fmt.Printf("%v is odd", a)
	}
}
