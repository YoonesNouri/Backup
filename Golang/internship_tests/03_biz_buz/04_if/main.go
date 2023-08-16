package main

import "fmt"

func main() {

	for i:=1 ; i <=100 ; i++ {
		if i%10 == 0 { fmt.Println(i,"buz")
		} else if i%5 == 0 { fmt.Println(i,"biz") }
	}
}
