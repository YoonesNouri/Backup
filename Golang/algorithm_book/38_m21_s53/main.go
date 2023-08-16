// اعداد دو رقمی که یکان و دهگان آنها فرد باشد
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Two-digit numbers with odd ones and tens:")

	for i := 10; i <= 99; i++ {
		if i%2 != 0 && (i/10)%2 != 0 {
			fmt.Println(i)
		}
	}

}
