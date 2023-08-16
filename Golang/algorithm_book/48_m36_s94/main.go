package main

import (
	"fmt"
	"math/rand"
)

var smallestindex int

func main() {
	numbers := []int{}
	for i := 0; i < 100; i++ {
		num := rand.Intn(200)
		numbers = append(numbers, num)
	}
	fmt.Println("numbers: ", numbers)

	smallest := numbers[0]
	for i, v := range numbers {
		if v < smallest {
			smallest = v
			smallestindex = i
		}
	}
	fmt.Println("smallest: ", smallest)
	fmt.Println("smallestindex: ", smallestindex)
}
