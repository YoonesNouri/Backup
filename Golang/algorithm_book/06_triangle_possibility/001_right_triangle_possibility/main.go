package main

import (
	"fmt"
	"math"
)

func isTriangleSide(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	return a+b > c && b+c > a && c+a > b && math.Pow(a, 2) == math.Pow(b, 2)+math.Pow(c, 2) || math.Pow(b, 2) == math.Pow(a, 2)+math.Pow(c, 2) || math.Pow(c, 2) == math.Pow(a, 2)+math.Pow(b, 2)
}

func main() {
	var a, b, c float64

	fmt.Print("Enter the lengths of three sides of a triangle (space-separated): ")
	fmt.Scanln(&a, &b, &c)

	if isTriangleSide(a, b, c) {
		fmt.Println("These lengths can form a right triangle.")

	} else {
		fmt.Println("These lengths cannot form a right triangle.")
	}
}
