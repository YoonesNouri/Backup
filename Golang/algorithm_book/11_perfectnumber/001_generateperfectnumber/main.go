package main

import "fmt"

func isPerfectNumber(num int) bool {
	sum := 0

	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num
}

func generatePerfectNumbers(limit int) []int {
	perfectNumbers := []int{}

	for i := 1; i <= limit; i++ {
		if isPerfectNumber(i) {
			perfectNumbers = append(perfectNumbers, i)
		}
	}

	return perfectNumbers
}

func main() {
	var limit int
	fmt.Print("enter limit:")
	fmt.Scanln(&limit)
	perfectNumbers := generatePerfectNumbers(limit)

	fmt.Println("Perfect numbers up to", limit, ":")
	for _, num := range perfectNumbers {
		fmt.Println(num)
	}
}

//package main

//import "fmt"

//func main() {
//	var limit int
//	fmt.Print("Enter the limit: ")
//	fmt.Scanln(&limit)

//	fmt.Println("Perfect numbers up to", limit, ":")
//	for num := 1; num <= limit; num++ {
//		sum := 0
//		for i := 1; i < num; i++ {
//			if num%i == 0 {
//				sum += i
//			}
//		}
//		if sum == num {
//			fmt.Println(num)
//		}
//	}
//}
