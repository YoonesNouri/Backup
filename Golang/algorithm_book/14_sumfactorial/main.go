package main //without functions

import "fmt"

func main() {
	var n, factorial, sumFactorials int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	factorial = 1
	sumFactorials = 0

	for i := 1; i <= n; i++ {
		factorial *= i
		sumFactorials += factorial
	}

	fmt.Println("Sum of factorials:", sumFactorials)
}

//package main //with functions

//import "fmt"

//func factorial(n int) int {
//	fac := 1
//	for i := 1; i <= n; i++ {
//		fac *= i
//	}
//	return fac
//}

//func sumfactorial(n int) int {
//	sumfac := 0
//	for i := 1; i <= n; i++ {
//		fac := factorial(i)
//		sumfac += fac
//	}
//	return sumfac
//}

//func main() {
//	var n int
//	fmt.Print("Enter a number: ")
//	fmt.Scanln(&n)

//	sf := sumfactorial(n)
//	fmt.Println("Sum of factorials:", sf)
//}
