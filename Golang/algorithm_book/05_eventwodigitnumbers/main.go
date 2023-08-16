package main

import (
	"fmt"
)

func main() {
	var s int = 0
	for i := 10; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			s += i
		}
	}
	fmt.Println("sum=", s)

}
