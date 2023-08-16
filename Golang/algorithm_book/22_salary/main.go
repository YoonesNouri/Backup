package main

import "fmt"

func calculateMonthlySalary(hoursWorked, hourlySalary float64) float64 {
	const monthlyWorkingHours = 160
	var monthlySalary float64

	if hoursWorked <= monthlyWorkingHours {
		monthlySalary = hoursWorked * hourlySalary
	} else {
		overtimeHours := hoursWorked - monthlyWorkingHours
		overtimePay := overtimeHours * (hourlySalary * 2 / 3)
		monthlySalary = (monthlyWorkingHours * hourlySalary) + overtimePay
	}

	return monthlySalary
}

func main() {
	var numEmployees int

	fmt.Print("Enter the number of employees: ")
	fmt.Scanln(&numEmployees)

	for i := 1; i <= numEmployees; i++ {
		var name string
		var hoursWorked, hourlySalary float64

		fmt.Printf("\nEmployee %d\n", i)

		fmt.Print("Enter employee name: ")
		fmt.Scanln(&name)

		fmt.Print("Enter total hours worked: ")
		fmt.Scanln(&hoursWorked)

		fmt.Print("Enter hourly salary: ")
		fmt.Scanln(&hourlySalary)

		monthlySalary := calculateMonthlySalary(hoursWorked, hourlySalary)
		fmt.Printf("Monthly salary of %s: $%.2f\n", name, monthlySalary)
	}
}
