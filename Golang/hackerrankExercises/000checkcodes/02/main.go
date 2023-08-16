package main

import "fmt"

func main() {
	a := []int32{-2, -1, 0, 1, 2}
	la := len(a)
	plusum := 0
	for i, ai := range a {

		if ai > 0 {
			plusum = ai[i] - ai[i-1]
			fmt.Println(i)

		}
	}
	tplusum := plusum
	fmt.Println(tplusum)

	per := float64(tplusum) / float64(la)
	fmt.Println(per*100, "%")
}
