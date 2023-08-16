package main

import "fmt"

var n, pn int
var sum float64

func main() {
	fmt.Print("enter a number greater than 2:")
	for {
		fmt.Scanln(&n)
		if n >= 2 {
			break
		}
		fmt.Print("please enter a number greater than 2:")
	}
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			pn = 1
		} else {
			pn = -1
		}
		sum += float64(pn) * float64(i-1) / float64(i)
	}
	fmt.Println("result:", sum)
}
