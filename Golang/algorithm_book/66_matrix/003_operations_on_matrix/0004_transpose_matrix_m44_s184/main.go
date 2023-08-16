// اگر تمام عناصر سطرها و ستونهای یک ماتریکس باهم عوض شود به ماترکس حاصل، ماتریکس ترانهاده گویند
// در واقع حول قطر اصلی ماتریکس ، قرینه میشود
// Transpose matrix
// https://fa.wikipedia.org/wiki/%D8%AA%D8%B1%D8%A7%D9%86%D9%87%D8%A7%D8%AF%D9%87
// Mr.c -> Mc.r
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

	// Create the empty matrix M
	M := make([][]int, rows)
	for i := 0; i < rows; i++ {
		M[i] = make([]int, cols)
	}

	// Taking values of matrix M
	fmt.Println("Enter matrix M values:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter value of M[%d][%d]: ", i+1, j+1)
			fmt.Scanln(&M[i][j])
		}
	}

	// Print matrix M
	fmt.Println("M = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", M[i][j])
		}
		fmt.Println()
	}

	// Create the empty matrix tM
	tM := make([][]int, rows)
	for i := 0; i < rows; i++ {
		tM[i] = make([]int, cols)
	}

	// replacing rows elements with columns elements.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tM[i][j] = M[j][i]
		}
	}

	// Print matrix tM
	fmt.Println("tM = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", tM[i][j])
		}
		fmt.Println()
	}

}
