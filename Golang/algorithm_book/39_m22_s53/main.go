// اعداد سه رقمی که یکان و هزارگان مساوی و دهگان آنها زوج باشد
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Three-digit numbers with equal ones and thousands and even tens:")

	for i := 100; i <= 999; i++ {
		ones := i % 10
		thousands := i / 100
		tens := (i / 10) % 10

		if ones == thousands && tens%2 == 0 {
			fmt.Println(i)
		}
	}
}
