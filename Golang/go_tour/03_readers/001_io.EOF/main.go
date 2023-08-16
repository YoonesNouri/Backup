package main

import (
	"fmt"
	"io"
)

func main() {
	var err error = io.EOF
	fmt.Println("err value: ",err) // Output: EOF
	fmt.Printf("err type: %T", err)
}
