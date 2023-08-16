package main

import (
	"fmt"
	"strconv"
)

func calculateNetSalary(grossSalary float64) float64 {
	insuranceDeduction := grossSalary * 0.10
	netSalary := grossSalary - insuranceDeduction
	return netSalary
}

func main() {
	fmt.Print("Enter the gross salary: ")
	var grossSalaryInput string
	_, err := fmt.Scanln(&grossSalaryInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	grossSalary, err := strconv.ParseFloat(grossSalaryInput, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	netSalary := calculateNetSalary(grossSalary)
	fmt.Printf("Net salary: %.2f\n", netSalary)
	//%.2f = با دقت دو رقم اعشار فرمت کن و نمایش بده
}
