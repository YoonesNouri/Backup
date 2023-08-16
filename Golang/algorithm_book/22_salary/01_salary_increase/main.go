package main

import "fmt"

func calculateSalaryincrease(salary float64) float64 {
	const level1 = 25
	const level2 = 35
	var increasedSalary float64

	if salary <= level1 {
		increasedSalary = salary * 1.05
	} else if salary > level1 && salary <= 35 {
		increasedSalary = salary * 1.07
	} else if salary > 35 {
		increasedSalary = salary * 1.10

	}

	return increasedSalary
}

func main() {
	var numEmployees int

	fmt.Print("Enter the number of employees: ")
	fmt.Scanln(&numEmployees)

	for i := 1; i <= numEmployees; i++ {
		var name string
		var salary, increasedSalary float64

		fmt.Printf("\nEmployee %d\n", i)

		fmt.Print("Enter employee name: ")
		fmt.Scanln(&name)

		fmt.Print("Enter employee's salary: ")
		fmt.Scanln(&salary)

		increasedSalary = calculateSalaryincrease(salary)
		fmt.Printf("increasedSalary salary of %s: $%.2f\n", name, increasedSalary)
	}
}
