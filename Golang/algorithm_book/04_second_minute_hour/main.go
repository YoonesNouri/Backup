package main

import (
	"fmt"
)

func calhourminutesecond(sec int) (h, m, s int) {
	h = sec / 3600
	m = (sec % 3600) / 60
	s = sec % 60
	fmt.Printf("It's %v hours, %v minutes, and %v seconds\n", h, m, s)
	return h, m, s
}

func main() {
	var sec int
	fmt.Print("Enter the number of seconds: ")
	fmt.Scanln(&sec)
	calhourminutesecond(sec)
}
