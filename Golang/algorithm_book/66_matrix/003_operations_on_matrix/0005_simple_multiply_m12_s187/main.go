// ضرب دو ماتریکس: باید ستون های اولی برابر با سطرهای دومی باشد A x B
// اگر تک تک عناصر سطرها در ستون هایی با عدد سطر و ستون برابر را در هم ضرب کنیم و جمع کنیم
// و سپس حاصلجمع را در نقطه تلاقی سطر و ستون جایگزین کنیم، ضرب ساده ی دو ماتریکس اتفاق افتاده
package main

import (
	"fmt"
)

func main() {
	var rowsA, colsA, rowsB, colsB int

	// Taking rows and columns quantities of matrix A.
	fmt.Println("<< Enter rows and columns of matrix A>>")
	fmt.Print("Enter rows: ")
	fmt.Scanln(&rowsA)
	fmt.Print("Enter columns: ")
	fmt.Scanln(&colsA)

	// Create the empty matrix A
	A := make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		A[i] = make([]int, colsA)
	}

	// Taking values of matrix A
	fmt.Println("Enter matrix A values:")
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			fmt.Printf("Enter value of A[%d][%d]: ", i+1, j+1)
			fmt.Scanln(&A[i][j])
		}
	}

	// Print matrix A
	fmt.Println("A = ")
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			fmt.Printf("%d  ", A[i][j])
		}
		fmt.Println()
	}
	//----------------------------------------------------------------------------
	// Taking rows and columns quantities of matrix B.
	fmt.Println("<< Enter rows and columns of matrix B>>")
	fmt.Print("Enter rows: ")
	fmt.Scanln(&rowsB)
	fmt.Print("Enter columns: ")
	fmt.Scanln(&colsB)

	// Create the empty matrix A
	B := make([][]int, rowsB)
	for i := 0; i < rowsB; i++ {
		B[i] = make([]int, colsB)
	}

	// Taking values of matrix B
	fmt.Println("Enter matrix B values:")
	for i := 0; i < rowsB; i++ {
		for j := 0; j < colsB; j++ {
			fmt.Printf("Enter value of B[%d][%d]: ", i+1, j+1)
			fmt.Scanln(&B[i][j])
		}
	}

	// Print matrix B
	fmt.Println("B = ")
	for i := 0; i < rowsB; i++ {
		for j := 0; j < colsB; j++ {
			fmt.Printf("%d  ", B[i][j])
		}
		fmt.Println()
	}
	//----------------------------------------------------------------------------
	// Create the empty matrix mulAB
	mulAB := make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		mulAB[i] = make([]int, colsB)
	}

	// A x B
	//در هر خانه از ماتریکس حاصلضرب به تعداد ستونهای اِی ، ضرب عناصر دریکدیگر هست که باهم جمع میشوند
	// ایندکس عناصر این ضربها تابع ایندکس ماتریکس حاصلضرب است
	//اولی ایندکس سطرش را از ایندکس سطر حاصلضرب میگیرد و دومی ایندکس ستونش را از ایندکس ستون حاصلضرب
	//این دو ایندکس ثابت اند و یک شمارنده ی دیگر بنام «کا» باید اضافه شود بتعداد ستونهای اِی تا بین ایندو ضرب کند
	//و در ایندکس حاصلضرب ، جمع آنها را ذخیره کند
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				mulAB[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	// Print matrix mulAB
	fmt.Println("mulAB = ")
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			fmt.Printf("%d  ", mulAB[i][j])
		}
		fmt.Println()
	}

}
