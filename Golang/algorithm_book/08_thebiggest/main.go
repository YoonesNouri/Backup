package main

import "fmt"

var a, b, c float64

func biggest(a, b, c float64) {
	if a > b && a > c {
		fmt.Printf("%v is the biggest number\n", a)
	} else if b > a && b > c {
		fmt.Printf("%v is the biggest number\n", b)
	} else {
		fmt.Printf("%v is the biggest number\n", c)
	}
}

func main() {
	fmt.Print("Enter three numbers: ")
	fmt.Scanln(&a, &b, &c)

	for a == b && b == c {
		fmt.Print("Enter three other numbers: ")
		fmt.Scanln(&a, &b, &c)
	}

	biggest(a, b, c)
}
