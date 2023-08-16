package main

import "fmt"

func main() {
	for num := 10; num < 100; num++ {
		tensdigit := num / 10  // Get the tens digit
		unitsdigit := num % 10 // Get the units digit
		inverse := unitsdigit*10 + tensdigit

		if num == inverse {
			fmt.Println(num)
		}
	}
}
