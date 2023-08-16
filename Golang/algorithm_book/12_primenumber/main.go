package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	divisors := []int{}

	//Find divisors of the number
	limit := int(math.Sqrt(float64(number)))
	for i := 1; i <= limit; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}

	//Check if the number of divisors is 2
	if len(divisors) == 1 {
		fmt.Printf("%d is a prime number.\n", number)
	} else {
		fmt.Printf("%d is not a prime number.\n", number)
	}
}

//for i := 2; i < int(math.Sqrt(float64(num))); i++ {
//		if num%i == 0 {
//	fmt.Println("it's not a prime number")
//		} else {fmt.Println("it's a prime number")
//}
