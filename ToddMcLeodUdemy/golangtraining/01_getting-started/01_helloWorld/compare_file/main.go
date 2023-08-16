package main

import (
	"fmt"
)

func isTriangleSide(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	return a+b > c && b+c > a && c+a > b
}

func isRightTriangle(a, b, c float64) bool {
	if !isTriangleSide(a, b, c) {
		return false
	}

	maxSide := max(max(a, b), c)

	// Check if the square of the longest side is equal to the sum of the squares of the other two sides
	if maxSide == a {
		return a*a == b*b+c*c
	} else if maxSide == b {
		return b*b == a*a+c*c
	} else {
		return c*c == a*a+b*b
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func drawRightTriangle() {
	fmt.Println("   *")
	fmt.Println("  **")
	fmt.Println(" ***")
	fmt.Println("****")
}

func main() {
	var a, b, c float64

	fmt.Print("Enter the lengths of three sides of a triangle (space-separated): ")
	fmt.Scanln(&a, &b, &c)

	if isRightTriangle(a, b, c) {
		fmt.Println("These lengths can form a right triangle.")
		drawRightTriangle()
	} else if isTriangleSide(a, b, c) {
		fmt.Println("These lengths can form a triangle, but it is not a right triangle.")
	} else {
		fmt.Println("These lengths cannot form a triangle.")
	}
}
