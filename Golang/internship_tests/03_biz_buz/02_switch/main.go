package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%10 == 0:
			fmt.Println(i, "buz")
		case i%5 == 0:
			fmt.Println(i, "biz")
		}
	}
}
