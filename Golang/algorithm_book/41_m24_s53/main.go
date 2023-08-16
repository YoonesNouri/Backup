// یک سکه پنجاه تومنی را به چند طریق میتوان با سکه های بیست، ده، و پنج تومنی خرد کرد؟
// (باید از همه ی سکهه ها استفاده شود)
package main

import (
	"fmt"
)

func main() {

	c50 := 50
	c20 := 20
	c10 := 10
	c5 := 5

	sp := []int{}
	for p20 := 1; p20 <= c50/c20; p20++ {
		for p10 := 1; p10 <= c50/c10; p10++ {
			for p5 := 1; p5 <= c50/c5; p5++ {
				if c50 == p20*c20+p10*c10+p5*c5 {
					sp = append(sp, p20, p10, p5)
				}
			}
		}
	}
	ways := len(sp) / 3
	fmt.Println("there are", ways, "ways")
}
