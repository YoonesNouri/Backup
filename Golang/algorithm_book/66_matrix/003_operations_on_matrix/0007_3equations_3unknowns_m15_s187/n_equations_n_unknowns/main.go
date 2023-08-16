package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {
	var n int

	// Prompt the user to enter the number of equations and unknowns.
	fmt.Print("Enter the number of equations and unknowns: ")
	fmt.Scanln(&n)

	// Create the coefficient matrix A with size n x n.
	A := mat.NewDense(n, n, nil)

	// Create the constants vector b with size n.
	b := mat.NewVecDense(n, nil)

	// Prompt the user to enter the values for the coefficient matrix A.
	fmt.Println("Enter the values for the coefficient matrix A:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var value float64
			fmt.Printf("Enter the value of A[%d][%d]: ", i+1, j+1)
			fmt.Scanln(&value)
			A.Set(i, j, value)
		}
	}

	// Prompt the user to enter the values for the constants vector b.
	fmt.Println("Enter the values for the constants vector b:")
	for i := 0; i < n; i++ {
		var value float64
		fmt.Printf("Enter the value of b[%d]: ", i+1)
		fmt.Scanln(&value)
		b.SetVec(i, value)
	}

	// Create a dense vector x with size n to store the solution.
	x := mat.NewVecDense(n, nil)

	// Solve the system of equations A * x = b.
	err := x.SolveVec(A, b)
	if err != nil {
		log.Fatal(err)
	}

	// Print the solution vector x.
	fmt.Println("Solution vector x:")
	for i := 0; i < n; i++ {
		fmt.Printf("x[%d] = %.2f\n", i+1, x.AtVec(i))
	}
}
