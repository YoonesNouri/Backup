//اعداد سه رقمی که با مقلوبشان برابر اند
package main

import (
	"fmt"
)

func RevSplitInt(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed
}

func main() {

for i:= 100 ; i<=999 ; i++ {
	if i == RevSplitInt(i) {
		fmt.Println(i)
	}
}
}