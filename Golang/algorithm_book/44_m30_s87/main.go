package main

import (
	"fmt"
)

var s, n int

func main() {
	for {
		fmt.Print("enter number of sentence(n): ")
		if _, err := fmt.Scanln(&n); err != nil {
			fmt.Println("invalid value.")
			continue
		}
		if n >= 0 {
			break
		}
		fmt.Println("invalid value.")
	}

	//1 , 2 , 3 , 4 , 5 i=شماره جمله
	//1 , 2 , 4 , 7 , 11 , 16 s=مقدار جمله
	//از جمله یک تا اِن اُم را پرینت کن
	s = 1
	for i := 1; i <= n; i++ {
		s = s + i - 1
		fmt.Printf("the %v-th sentence: %v\n", i, s)
	}
}
