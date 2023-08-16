// اسلایسی بساز بطول 20 که ایندکس های فرد مقادیر 1 تا 10 و ایندکس های زوج مقادیر 11 تا 20 باشد
package main

import (
	"fmt"
)

func main() {
	n := 20
	slice := make([]int, n)

	for i := 0; i < n/2; i++ {
		slice[2*i+1] = i
		slice[2*i] = i + (n / 2)
	}
	fmt.Println("slice:", slice)
}
