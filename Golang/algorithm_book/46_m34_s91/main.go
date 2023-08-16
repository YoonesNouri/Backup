package main

import (
	"fmt"
)

func main() {
capital := map[string]string{
		"iran": "tehran",
		"iraq": "baqdad",
		"china": "Beijing",
	}	
fmt.Println("the capital of iran is",capital["iran"])
}