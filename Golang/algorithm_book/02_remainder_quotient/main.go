package main

import (
	"fmt"
	"math"
)

var a, b float64

func remainder(a, b float64) float64 {
	return a - (b * (math.Floor(a / b)))
}

func quotient(a, b float64) float64 {
	return (a / b) - ((remainder(a, b)) / b)
}

func main() {

	//fmt.Println("remainder=", remainder(5, 2))
	fmt.Println("quotient=", quotient(5, 2))

}
