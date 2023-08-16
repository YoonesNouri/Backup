// اعداد چهار رقمی که یکان و صدگان زوج اما دهگان و هزارگان فرد باشد
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Four-digit numbers with even ones and hundreds but odd tens and thousands:")

	for i := 1000; i <= 9999; i++ {
		if i%2 == 0 && (i/100)%2 == 0 && (i/10)%2 != 0 && (i/1000)%2 != 0 {
			fmt.Println(i)
		}
	}

}
