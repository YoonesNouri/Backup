package main

import (
	"fmt"
)

var limit int
var sumlog float64

func fac(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func naturalLogarithm(limit int) float64 {
	sumlog = 1
	for i := 1; i <= limit; i++ {
		sumlog += 1 / float64(fac(i))
	}
	return sumlog
}

func main() {
	fmt.Print("enter limit:")
	fmt.Scanln(&limit)

	Nlog := naturalLogarithm(limit)

	fmt.Println("Natural Logarithm number approaches: ", Nlog)

}
