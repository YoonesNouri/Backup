// فرمول زیر را در اندیکس های یک اسلایس جایگذاری کن
package main

import (
	"fmt"
)

func main() {
	n := 10
	A := make([]int, n)

	for i := 1; i < n; i++ {
		A[0] = 2
		A[i] = 2*A[i-1] + 3*i
	}
	fmt.Println("A:", A)
}
