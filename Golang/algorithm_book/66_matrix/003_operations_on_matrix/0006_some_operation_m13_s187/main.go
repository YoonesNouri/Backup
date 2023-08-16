// ماتریکسی که خانه اولش 1 و مابقی خانه هایش حاصلجمع ایندکس همان خانه باشد
// توجه: ایندکس ماتریکس در ریاضی از 1 شروع میشود، نه مثل گولنگ از 0
package main

import (
	"fmt"
)

func main() {
	var rowscols int

	// Taking rows and columns quantities of matrix A.
	fmt.Println("<< Enter rows and columns of matrix A>>")
	fmt.Print("Enter rowscols: ")
	fmt.Scanln(&rowscols)

	// Create the empty matrix A
	A := make([][]int, rowscols)
	for i := 0; i < rowscols; i++ {
		A[i] = make([]int, rowscols)
	}

	// Taking values of matrix A
	fmt.Println("Enter matrix A values:")
	for i := 0; i < rowscols; i++ {
		for j := 0; j < rowscols; j++ {
			if i == 0 && j == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = i + 1 + j + 1
			}
		}
	}

	// Print matrix A
	fmt.Println("A = ")
	for i := 0; i < rowscols; i++ {
		for j := 0; j < rowscols; j++ {
			fmt.Printf("%d  ", A[i][j])
		}
		fmt.Println()
	}

}
