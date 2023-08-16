package main

import (
	"fmt"
	"math"
)

var radius int

func main() {
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanln(&radius)

	fmt.Printf("The intersection points of a circle with a radius of %v at the center of coordinates with integers:\n", radius)

	for x := -radius; x <= radius; x++ {
		ySquare := radius*radius - x*x
		if ySquare >= 0 {
			y := int(math.Sqrt(float64(ySquare)))
			if y*y == ySquare {
				fmt.Printf("(%d, %d)\n", x, y)
				if y != 0 {
					fmt.Printf("(%d, %d)\n", x, -y) // Additional point reflected across the x-axis
				}
			}
		}
	}
}
