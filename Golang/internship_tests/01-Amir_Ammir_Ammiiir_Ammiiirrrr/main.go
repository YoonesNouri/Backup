package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	//fmt.Scanln(&input)
	input = "Amir"
	spl := strings.Split(input, "")

	for i := 0; i < len(spl); i++ {
		out := spl[:i+1]

		rep := ""
		for index, value := range out {
			rep += strings.Repeat(value, index+1)
		}

		j := strings.Join(spl[i+1:], "")

		fmt.Println("rep =", rep)
		fmt.Println("j =", j)
		fmt.Println("rep + j =", rep + j)

	}
}
