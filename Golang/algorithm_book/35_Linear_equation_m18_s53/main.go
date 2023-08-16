package main

import (
	"fmt"
)

func main() {
	var dx, dy, x1, x2, y1, y2 float64

	fmt.Print("Enter the coordinates (longitude, latitude) of point 1:\n")
	fmt.Print("Enter x1: ")
	for {
		if _, err := fmt.Scanln(&x1); err != nil {
			fmt.Println("Invalid input. Please enter a valid numeric value.")
			continue
		}

		break
	}

	fmt.Print("Enter y1: ")
	for {
		if _, err := fmt.Scanln(&y1); err != nil {
			fmt.Println("Invalid input. Please enter a valid numeric value.")
			continue
		}

		break
	}

	p1 := []float64{x1, y1}

	fmt.Print("Enter the coordinates (longitude, latitude) of point 2:\n")
	fmt.Print("Enter x2: ")
	for {
		if _, err := fmt.Scanln(&x2); err != nil {
			fmt.Println("Invalid input. Please enter a valid numeric value.")
			continue
		}

		break
	}

	fmt.Print("Enter y2: ")
	for {
		if _, err := fmt.Scanln(&y2); err != nil {
			fmt.Println("Invalid input. Please enter a valid numeric value.")
			continue
		}

		break
	}

	p2 := []float64{x2, y2}

	//دلتا ایکس و دلتاایگرگ و اِم برای محاسبه ی شیب خط
	//dx = x2 - x1
	//dy = y2 - y1
	dx = p2[0] - p1[0]
	dy = p2[1] - p1[1]
	m := dy / dx

	//یک نقطه رادر معادله جایگزین میکنیم تا عرض از مبدا بدست بیاید x,y
	// y = mx + b --> b = y - mx
	b := p1[1] - m*p1[0]

	mString := formatCoefficient(m)
	bString := formatIntercept(b)

	fmt.Printf("The equation of the line is y = %s x %s\n", mString, bString)
}

func formatCoefficient(m float64) string {
	if m == 1 {
		return ""
	} else if m == -1 {
		return "-"
	}
	return fmt.Sprintf("%v", m)
}

func formatIntercept(b float64) string {
	if b != 0 {
		if b < 0 {
			return fmt.Sprintf("- %v", -b)
		} else {
			return fmt.Sprintf("+ %v", b)
		}
	}
	return ""
}
