package main

import (
	"fmt"
)

func main() {
	var rows, cols int

	// Taking rows and columns quantities.
	fmt.Println("<< Enter rows and columns of matrix >>")
	fmt.Print("Enter rows: ")
	fmt.Scanln(&rows)
	fmt.Print("Enter columns: ")
	fmt.Scanln(&cols)

	// Create the matrix A
	A := make([][]int, rows)
	for i := 0; i < rows; i++ {
		A[i] = make([]int, cols)
	}

	// Taking values of matrix A
	fmt.Println("Enter matrix A values:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter value of A[%d][%d]: ", i+1, j+1)
			fmt.Scanln(&A[i][j])
		}
	}
	// Print matrix
	fmt.Println("A = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", A[i][j])
		}
		fmt.Println()
	}

	// Taking the row number to calculate the sum
	var rownum int
	fmt.Print("Enter the row number to calculate the sum: ")
	fmt.Scanln(&rownum)

	// Calculate the sum of the specified row
	sumrow := 0
	for _, v := range A[rownum-1] {
		sumrow += v
	}

	// Print the sum
	fmt.Printf("Sum of A[%d]: %d\n", rownum, sumrow)

	// Taking the column number to calculate the sum
	var colnum int
	fmt.Print("Enter the column number to calculate the sum: ")
	fmt.Scanln(&colnum)

	// Calculate the sum of the specified column
	sumcol := 0
	for i := 0; i < rows; i++ {
		sumcol += A[i][colnum-1]
	}

	// Print the sum
	fmt.Printf("Sum of A[][%d]: %d\n", colnum, sumcol)

}
