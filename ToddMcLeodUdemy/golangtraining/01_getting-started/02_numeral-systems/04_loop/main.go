package main

import "fmt"

func main() {
	for i := 1000; i < 1100; i++ {
		fmt.Printf("%d \t %q \t %b \t %x \n", i, i, i, i)
	}
}
