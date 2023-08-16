package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n1 := 10
	RNDlimit1 := 100
	numbers1 := []int{}
	for i := 0; i < n1; i++ {
		num := rand.Intn(RNDlimit1)
		numbers1 = append(numbers1, num)
	}
	fmt.Println("numbers1:", numbers1)

	n2 := 10
	RNDlimit2 := 100
	numbers2 := []int{}
	for i := 0; i < n2; i++ {
		num := rand.Intn(RNDlimit2)
		numbers2 = append(numbers2, num)
	}
	fmt.Println("numbers2:", numbers2)

	// Determine the smaller limit for unequal length slices
	n := n1
	if n2 < n1 {
		n = n2
	}

	// numbers1 + numbers2
	sumslice := make([]int, n)
	for i := 0; i < n1; i++ {
		sumslice[i] = numbers1[i] + numbers2[i]
	}
	fmt.Println("sumslice:", sumslice)

	// numbers1 - numbers2
	subslice := make([]int, n)
	for i := 0; i < n1; i++ {
		subslice[i] = numbers1[i] - numbers2[i]
	}
	fmt.Println("subslice:", subslice)

	// numbers1 * numbers2
	mulslice := make([]int, n)
	for i := 0; i < n1; i++ {
		mulslice[i] = numbers1[i] * numbers2[i]
	}
	fmt.Println("mulslice:", mulslice)

	//sum of mulslice elements
	summulslice := 0
	for _, v := range mulslice {
		summulslice += v
	}
	fmt.Println("summulslice:", summulslice)

}
