package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy) // dy = length of the outer slice of a
	for x := range a {       // ranges over outer slice a and
		a[x] = make([]uint8, dx) //initializes each inner slice with a length of dx
	}
	for x := range a { //دو تا رینچ تو دل هم
		for y := range a[x] { // در نقطه ایکس(که خودش یک اسلایس یک بعدی است)، ایگرگ چند است؟
			a[y][x] = uint8(x + y) //به مختصات نقطه ی ایکس ایگرگ، ایکس+ایگرگ را اختصاص بده
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
