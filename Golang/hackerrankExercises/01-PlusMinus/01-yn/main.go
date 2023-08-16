package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, -3, -1, 9, 1, 3, 0, 0, -4, 4}
	sort.Ints(a)
	la := len(a)

	minsum := 0
	for i, ai := range a {

		if ai < 0 {
			minsum += i

		}
	}

	permin := float64(minsum) / float64(la)
	fmt.Println("total numbers = ", la, " , minus numbers = ", minsum, " , percentage = ", permin*100, "%")

	possum := 0
	for i, ai := range a {

		if ai > 0 {
			possum += (i - (i - 1))

		}
	}

	perpos := float64(possum) / float64(la)
	fmt.Println("total numbers = ", la, " , positive numbers = ", possum, " , percentage = ", perpos*100, "%")

	plszero := 0
	for i, ai := range a {

		if ai == 0 {
			plszero += (i - (i - 1))

		}
	}

	perzero := float64(plszero) / float64(la)
	fmt.Println("total numbers = ", la, " , zero numbers = ", plszero, " , percentage = ", perzero*100, "%")
}
