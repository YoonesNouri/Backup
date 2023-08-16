//اعداد سه رقمی ، همه ی رقم ها زوج
package main

import (
"fmt"
)

func main() {
	for i := 100 ; i <=999 ; i++{
		if (i/100)%2 == 0 && (i/10)%2 == 0 && i%2 == 0 {
			fmt.Println(i)
		}
	}
}