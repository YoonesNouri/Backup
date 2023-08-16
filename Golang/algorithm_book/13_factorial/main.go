package main

import (
	"fmt"
)

func Fac(n int) { // removed int from return 
	fc := 1
	for i := 1; i <= n; i++ {
		fc *= i
	}
	//return fc
	fmt.Println(fc)
}

func main() {
//	fac := Fac(4)
//fmt.Println(fac)
Fac(4)
}
