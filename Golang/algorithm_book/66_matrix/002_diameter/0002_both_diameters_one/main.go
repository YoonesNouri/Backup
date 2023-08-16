// دو قطر اسلایس مربعی 1 و باقی ایندکس هایش صفر باشد
package main

import (
	"fmt"
)

var rows, cols int

func main() {
	// taking rows and columns values.
	fmt.Print("<< Enter rows and columns of matrix >> \n")
	fmt.Print("Enter rows : ")
	fmt.Scanln(&rows)
	fmt.Print("Enter columns : ")
	fmt.Scanln(&cols)

	//making empty matrix
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	//Assigning values to indexes on the both diameters.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == j || i+j == rows-1 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}

	// Print matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", matrix[i][j])
		}
		fmt.Println()
	}

}
