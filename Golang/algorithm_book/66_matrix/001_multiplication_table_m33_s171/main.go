//جدول ضرب 10 در 10 با اسلایس دو بعدی
package main

import (
	"fmt"
)

func main() {
	twoDslice := make([][]int, 10)
	for row := 0; row < 10; row++ {
		twoDslice[row] = make([]int, 10) //allocate memory for each row before initializing.
		for column := 0; column < 10; column++ {
			twoDslice[row][column] = (row + 1) * (column + 1)
		}
	}

// Print the multiplication table
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d\t", twoDslice[i][j])
		}
		fmt.Println()
	}
}
