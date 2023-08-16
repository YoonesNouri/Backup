package main

import "fmt"

func main() {
	for i := 5; i <= 100; i += 5 {
		message := "biz"
		if i%10 == 0 {
			message = "buz"
		}
		fmt.Println(i, message)
	}
}
