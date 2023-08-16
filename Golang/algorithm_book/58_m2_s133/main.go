// سه عدد را بگیر و به ترتیب معکوس برگردان
package main

import (
	"fmt"
)

var a int

func main() {
	sa := []int{}
	for {
		fmt.Print("Enter a number : ")
		_, err := fmt.Scanln(&a)
		if err != nil {
			break
		}
		sa = append(sa, a)
	}
	rsa := []int{}
	for i := len(sa) - 1; i >= 0; i-- {
		rsa = append(rsa, sa[i])
	}
	fmt.Println(sa)
	fmt.Println(rsa)
}
