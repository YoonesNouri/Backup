package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&num)
	count := 0
	for {
		num /= 10
		count++
		if num == 0 {
			break
		}
	}
	fmt.Println("digits: ", count)
}
