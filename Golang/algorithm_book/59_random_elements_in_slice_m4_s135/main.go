package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 100
	RNlimit := 200

	numbers := []int{}
	for i := 0; i < n; i++ {
		num := rand.Intn(RNlimit)
		numbers = append(numbers, num)
	}
	fmt.Println("numbers: ", numbers)

	sum := 0
	for _, v := range numbers {
		sum += v
	}
	ave := sum / len(numbers)
	fmt.Println("average:", ave)

	smallerThanAverage := []int{}
	biggerThanAverage := []int{}
	for _, v := range numbers {
		if v < ave {
			smallerThanAverage = append(smallerThanAverage, v)
		} else if v > ave {
			biggerThanAverage = append(biggerThanAverage, v)
		}
	}

	fmt.Printf("%v numbers are smaller than average\n", len(smallerThanAverage))
	fmt.Printf("%v numbers are bigger than average", len(biggerThanAverage))
}
