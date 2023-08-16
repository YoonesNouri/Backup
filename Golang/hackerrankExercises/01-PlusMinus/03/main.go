package main

import (
	"fmt"
)

func main() {

	var positive, negative, zero, size float64

	size = int(size)
	for v := 0; v < size; v++ {
		if v > 0 {
			positive++
		} else if v < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%.6f\n", positive/size)
	fmt.Printf("%.6f\n", negative/size)
	fmt.Printf("%.6f\n", zero/size)
}
