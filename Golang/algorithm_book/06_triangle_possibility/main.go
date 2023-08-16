package main

import "fmt"

func isTriangleSide(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	return a+b > c && b+c > a && c+a > b

}

func main() {
	var a, b, c float64

	fmt.Print("Enter the lengths of three sides of a triangle (space-separated): ")
	fmt.Scanln(&a, &b, &c)

	if isTriangleSide(a, b, c) {
		fmt.Println("These lengths can form a triangle.")

	} else {
		fmt.Println("These lengths cannot form a triangle.")
	}
}
