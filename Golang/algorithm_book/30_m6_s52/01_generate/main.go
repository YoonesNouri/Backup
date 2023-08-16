package main

import (
	"fmt"
)

var limit int

func sumDigits(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit
		num = num / 10
	}
	return sum
}

func isDivisible(limit, sumdigit int) bool {
	if limit%sumdigit == 0 {
		return true
	} else {
		return false
	}
}

func generateDivisible(limit int) []int {
	divisibleNumbers := []int{}
	for i := 1; i <= limit; i++ {
		if isDivisible(i, sumDigits(i)) {
			divisibleNumbers = append(divisibleNumbers, i)
		}
	}
	return divisibleNumbers
}

func main() {
	fmt.Print("Enter limit (limit>0): ")
	for {
		fmt.Scanln(&limit)
		if limit > 0 {
			break
		}
		fmt.Print("Enter limit (limit>0): ")
	}

	fmt.Println("Generated Divisible Numbers:")
	divisibleNumbers := generateDivisible(limit)
	fmt.Println(divisibleNumbers)
}

//ساده ترین کدش اینه
//package main

//import (
//	"fmt"
//)

//func sumDigits(num int) int {
//	sum := 0
//	for num != 0 {
//		digit := num % 10
//		sum += digit
//		num /= 10
//	}
//	return sum
//}

//func main() {
//	var limit int
//	fmt.Print("Enter limit (limit>0): ")
//	for {
//		fmt.Scanln(&limit)
//		if limit > 0 {
//			break
//		}
//		fmt.Print("Enter limit (limit>0): ")
//	}

//	fmt.Println("Generated Divisible Numbers:")
//	for i := 1; i <= limit; i++ {
//		if i%sumDigits(i) == 0 {
//			fmt.Println(i)
//		}
//	}
//}
