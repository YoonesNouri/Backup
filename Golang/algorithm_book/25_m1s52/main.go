package main

import "fmt"

var n1, n2, n3, n4 int

func main() {
	fmt.Print("enter 4 numbers:\n")
	fmt.Print("enter number 1: ")
	fmt.Scanln(&n1)
	fmt.Print("enter number 2: ")
	fmt.Scanln(&n2)
	fmt.Print("enter number 3: ")
	fmt.Scanln(&n3)
	fmt.Print("enter number 4: ")
	fmt.Scanln(&n4)

	if n1 > n3+n4 {
		fmt.Println("n1 x n3 = ", n1*n3)
	} else {
		fmt.Println("n2 x n4 = ", n2*n4)
	}

}
