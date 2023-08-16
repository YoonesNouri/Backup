package main

import (
	"fmt"
)

func calculateFine(violationCode, numViolations int) int {
	// Define the fine amounts for each violation code
	fineAmounts := map[int]int{
		1:  100, // Fine amount for violation code 1
		2:  200, // Fine amount for violation code 2
		3:  150, // Fine amount for violation code 3
		4:  250, // Fine amount for violation code 4
		5:  180, // Fine amount for violation code 5
		6:  300, // Fine amount for violation code 6
		7:  220, // Fine amount for violation code 7
		8:  270, // Fine amount for violation code 8
		9:  190, // Fine amount for violation code 9
		10: 350, // Fine amount for violation code 10
	}

	fine, exists := fineAmounts[violationCode]
	if !exists {
		return 0 // Invalid violation code, return 0 fine amount
	}

	return fine * numViolations
}

func main() {
	var violationCode, numViolations int

	// Prompt the user to enter the violation code and number of violations
	fmt.Print("Enter the violation code: ")
	fmt.Scanln(&violationCode)
	fmt.Print("Enter the number of violations: ")
	fmt.Scanln(&numViolations)

	totalFine := calculateFine(violationCode, numViolations)
	fmt.Println("Total fine amount:", totalFine)
}
