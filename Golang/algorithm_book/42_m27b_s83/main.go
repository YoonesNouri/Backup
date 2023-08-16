package main

import (
	"fmt"
	"math"

	"github.com/YoonesNouri/mymodule/mymath"
)

var n int
var x float64

func main() {
	fmt.Print("Enter n:")
	fmt.Scanln(&n)
	fmt.Print("Enter x:")
	fmt.Scanln(&x)

	// sum = 1 - x^2/2! + ... + (-1)^(2n+1) * x^2n/(2n)!
	sum := 1.0
	for i := 1; i <= n; i++ {
		sum += math.Pow(-1, float64(2*i+1)) * math.Pow(x, float64(2*i)) / float64(mymath.Fac(2*i))
	}
	fmt.Println("sum =", sum)
}
