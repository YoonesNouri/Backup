// جواب دستگاه سه معادله سه مجهول را به وسیله ی سه ماتریکس بدست بیاورید
// با استفاده از این بستر
// A[0][0]*x + A[0][1]*y + A[0][2]*z = b[0]
// A[1][0]*x + A[1][1]*y + A[1][2]*z = b[1]
// A[2][0]*x + A[2][1]*y + A[2][2]*z = b[2]
package main

import (
	"fmt"
	"math"
)

func main() {
	// Coefficient matrix A
	A := [][]float64{
		{2, 1, -1},
		{-3, -1, 2},
		{-2, 1, 2},
	}

	// Constants matrix B
	B := make([]float64, 3)
	fmt.Println("Enter values for the constants matrix B:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Enter value of B[%d]: ", i)
		fmt.Scanln(&B[i])
	}

	// Calculate the inverse of matrix A
	AInverse, err := inverseMatrix(A)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calculate X = A^(-1) * B
	X, err := multiplyMatrixVector(AInverse, B)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the solution
	fmt.Println("Solution:")
	fmt.Printf("x = %.2f\n", X[0])
	fmt.Printf("y = %.2f\n", X[1])
	fmt.Printf("z = %.2f\n", X[2])
}

//--------------------------------------------------------------------------------------------------

// Calculates the inverse of a 3x3 matrix
func inverseMatrix(A [][]float64) ([][]float64, error) {
	det := determinant(A)
	if math.Abs(det) < 1e-8 {
		return nil, fmt.Errorf("matrix is not invertible")
	}

	adj := adjugate(A)
	inv := scalarMultiplyMatrix(adj, 1/det)
	return inv, nil
}

// Calculates the determinant of a 3x3 matrix
func determinant(A [][]float64) float64 {
	a, b, c := A[0][0], A[0][1], A[0][2]
	d, e, f := A[1][0], A[1][1], A[1][2]
	g, h, i := A[2][0], A[2][1], A[2][2]

	return a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
}

// Calculates the adjugate of a 3x3 matrix
func adjugate(A [][]float64) [][]float64 {
	a, b, c := A[0][0], A[0][1], A[0][2]
	d, e, f := A[1][0], A[1][1], A[1][2]
	g, h, i := A[2][0], A[2][1], A[2][2]

	return [][]float64{
		{e*i - f*h, c*h - b*i, b*f - c*e},
		{f*g - d*i, a*i - c*g, c*d - a*f},
		{d*h - e*g, b*g - a*h, a*e - b*d},
	}
}

// Multiplies a matrix by a scalar
func scalarMultiplyMatrix(A [][]float64, scalar float64) [][]float64 {
	result := make([][]float64, len(A))
	for i := 0; i < len(A); i++ {
		result[i] = make([]float64, len(A[i]))
		for j := 0; j < len(A[i]); j++ {
			result[i][j] = A[i][j] * scalar
		}
	}
	return result
}

// Multiplies a matrix by a vector
func multiplyMatrixVector(A [][]float64, v []float64) ([]float64, error) {
	if len(A[0]) != len(v) {
		return nil, fmt.Errorf("matrix and vector dimensions don't match")
	}

	result := make([]float64, len(A))
	for i := 0; i < len(A); i++ {
		sum := 0.0
		for j := 0; j < len(A[i]); j++ {
			sum += A[i][j] * v[j]
		}
		result[i] = sum
	}
	return result, nil
}
