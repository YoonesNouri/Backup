package main

import (
	"fmt"
)

var a, b int
var common3 []int

func main() {

	fmt.Print("Enter 2 numbers (a,b>0):\n")
	for {
		fmt.Print("Enter numbers 1 :")
		fmt.Scanln(&a)
		fmt.Print("Enter numbers 2 :")
		fmt.Scanln(&b)

		if a > 0 && b > 0 {
			break
		}
		fmt.Print("Enter 2 numbers (a,b>0):")
	}

	if a < b {
		a, b = b, a
	}
	// a is always the bigger numbers and b is the smaller.
	// The common factor of 3 for two numbers
	//is actually the factor of 3 for the smaller number within its range.
	//مضارب ۳ مشترک برای دو عدد، در واقع مضارب ۳ برای عدد کوچکتر است

	for i := 1; i <= b; i++ {
		if i%3 == 0 && a%i == 0 && b%i == 0 {
			common3 = append(common3, i)
		}
	}
	fmt.Printf("The common factor of 3 for %v and %v : %v", a, b, common3)
}
