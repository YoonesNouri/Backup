package main

import (
	"fmt"
	"math"
)

var n int
var s float64

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

	//1 , 2 , 3 , 4 شماره جمله = i
	//1,-1/3,1/5,-1/7 مقدار جمله = s
	//تا جمله اِنُم و مجموع جملات را پرینت کن 
	sum := 0.0
	for i := 1; i <= n; i++ {
		s = math.Pow(-1, float64(i+1)) * float64(1/float64(2*i - 1))
		fmt.Printf("the %v-th sentence: %v\n", i, s)
		sum += s
	}
	fmt.Printf("sum of sentences till %v-th : %v", n, sum)
}