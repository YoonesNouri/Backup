// عوض کردن عناصر دو سطر یا دو ستون  با هم
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

	// Create the matrix S
	S := make([][]int, rows)
	for i := 0; i < rows; i++ {
		S[i] = make([]int, cols)
	}

	// Taking values of matrix S
	fmt.Println("Enter matrix S values:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter value of S[%d][%d]: ", i+1, j+1)
			fmt.Scanln(&S[i][j])
		}
	}

	// Print matrix S
	fmt.Println("S = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", S[i][j])
		}
		fmt.Println()
	}

	// Taking the number of two rows to be replaced
	var rownumA, rownumB int
	fmt.Print("Enter the number of two rows to be replaced: \n")
	fmt.Print("rownumA: ")
	fmt.Scanln(&rownumA)
	fmt.Print("rownumB: ")
	fmt.Scanln(&rownumB)

	// replacing rowA with rowB
	rowtransS := S
	for j := 0; j < cols; j++ {
		rowtransS[rownumA-1][j], rowtransS[rownumB-1][j] = rowtransS[rownumB-1][j], rowtransS[rownumA-1][j]
	}

	// Print matrix rowtransS
	fmt.Println("rowtransS = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", rowtransS[i][j])
		}
		fmt.Println()
	}

	// Taking the number of two columns to be replaced
	var colnumA, colnumB int
	fmt.Print("Enter the number of two columns to be replaced: \n")
	fmt.Print("colnumA: ")
	fmt.Scanln(&colnumA)
	fmt.Print("colnumB: ")
	fmt.Scanln(&colnumB)

	// replacing columnA with columnB
	coltransS := S
	for i := 0; i < rows; i++ {
		coltransS[i][colnumA-1], coltransS[i][colnumB-1] = coltransS[i][colnumB-1], coltransS[i][colnumA-1]
	}

	// Print matrix coltransS
	fmt.Println("coltransS = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", coltransS[i][j])
		}
		fmt.Println()
	}

}
