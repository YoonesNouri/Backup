package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 int

	fmt.Print("Enter 3 unequal numbers:\n")

	for {
		fmt.Print("Enter number 1:")
		fmt.Scanln(&n1)
		fmt.Print("Enter number 2:")
		fmt.Scanln(&n2)
		fmt.Print("Enter number 3:")
		fmt.Scanln(&n3)

		if !(n1 == n2 || n1 == n3 || n2 == n3) {
			break
		}
		fmt.Print("please Enter 3 unequal numbers:\n")
	}

	fmt.Println("sorting from small to big...")

	numbers := []int{n1, n2, n3}
	sort.Ints(numbers)
	//sort.Ints function does not return the sorted slice.
	//It sorts the slice in-place.
	//so you don't need to assign the result to a variable like this : sorted := sort.Ints(numbers).

	fmt.Println(numbers)
	fmt.Printf("or can be printed like this: ")
	fmt.Printf("%v < %v < %v", numbers[0], numbers[1], numbers[2])

}
