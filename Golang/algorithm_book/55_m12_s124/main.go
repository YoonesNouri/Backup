// √1 + x√1 + x == √(1 + x)*√(1 + x)) :محاسبه دنباله که من منظورش را این فرض کردم
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var x, result float64

	fmt.Print("enter x: ")
	fmt.Scanln(&x)
	fmt.Print("enter n: ")
	fmt.Scanln(&n)

	result = math.Sqrt(1 + x)
	for i := 1; i <= n; i++ {
		result = math.Sqrt((1 + x) * result)
	}
	fmt.Println("result = ", result)
}
