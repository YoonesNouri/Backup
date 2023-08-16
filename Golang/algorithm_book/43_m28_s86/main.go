package main

import (
	"fmt"
	"math"
)

var n int

func sqrt3(n int) float64 {
	if n == 0 {
		return math.Sqrt(3)
	}
	return math.Sqrt(3 + sqrt3(n-1))
}

func main() {
	fmt.Print("Enter the number of nested radicals (n >= 0): ")
	for {
		if _, err := fmt.Scanln(&n); err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		if n >= 0 {
			break
		}
		fmt.Println("Please enter a positive number (n >= 0).")
	}

	sum := sqrt3(n)
	fmt.Printf("The sum of the nested radicals of 3 up to the %d-th term is: %f\n", n, sum)
}

//package main

//import (
//	"fmt"
//	"math"
//)
//میشه این فانکشنو حذف کرد و فقط لوپ بمونه
//func sqrtSeries(n int) float64 {
//	sum := math.Sqrt(3)
//	for i := 1; i < n; i++ {
//		sum = math.Sqrt(3 + sum)
//	}
//	return sum
//}

//func main() {
//	var n int
//	fmt.Print("Enter the number of nested radicals (n > 0): ")
//	for {
//		if _, err := fmt.Scanln(&n); err != nil {
//			fmt.Println("Invalid input. Please try again.")
//			continue
//		}
//		if n > 0 {
//			break
//		}
//		fmt.Println("Please enter a positive number (n > 0).")
//	}

//	sum := sqrtSeries(n)
//	fmt.Printf("The sum of the nested radicals of 3 up to the %d-th term is: %f\n", n, sum)
//}
