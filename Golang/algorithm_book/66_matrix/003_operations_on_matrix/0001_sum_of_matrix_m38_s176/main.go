// دو ماتریکس آ و بی را باهم جمع کن و در ماتریکس سی تعبیه کن
// تفریق و ضرب و تقسیم هم همین کد است فقط جای علامت بعلاوه با اینها عوض میشه
package main

import (
	"fmt"
)

var rows, cols int

func main() {
	// taking rows and columns quantities.
	fmt.Print("<< Enter rows and columns of matrix >> \n")
	fmt.Print("Enter rows : ")
	fmt.Scanln(&rows)
	fmt.Print("Enter columns : ")
	fmt.Scanln(&cols)

	//making matrix A
	A := make([][]int, rows)
	for i := 0; i < rows; i++ {
		A[i] = make([]int, cols)
	}
	//taking values of matrix A
	fmt.Println("Enter matrix A values:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter value of A[%d][%d] : ", i+1, j+1)
			_, err := fmt.Scan(&A[i][j])
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
		}
	}

	//making matrix B
	B := make([][]int, rows)
	for i := 0; i < rows; i++ {
		B[i] = make([]int, cols)
	}
	//taking values of matrix B
	fmt.Println("Enter matrix B values:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter value of B[%d][%d] : ", i+1, j+1)
			_, err := fmt.Scan(&B[i][j])
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
		}
	}

	//making matrix C
	C := make([][]int, rows)
	for i := 0; i < rows; i++ {
		C[i] = make([]int, cols)
	}
	//C = A + B
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}

	// Print matrix
	fmt.Println("C = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", C[i][j])
		}
		fmt.Println()
	}
}
