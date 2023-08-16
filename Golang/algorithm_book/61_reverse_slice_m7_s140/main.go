package main

import (
	"fmt"
)

var num int

func main() {
	slice := []int{}
	for {
		fmt.Print("enter a number: ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		slice = append(slice, num)
	}
	fmt.Println("slice:        ", slice)

	rs := slice
	j := len(rs) - 1
	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[j] = rs[j], rs[i]
		j--
	}
	fmt.Println("reverse slice:", rs)

}
