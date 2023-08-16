// المنت های مشترک دو اسلایس با طول مساوی که المنت تکراری ندارند را در یک اسلایس تعبیه کن
package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("make two slices :")

	slice1 := []int{}
	for {
		fmt.Print("Enter elements of slice1: ")
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}
		slice1 = append(slice1, n)
	}
	fmt.Println("slice1 :", slice1)

	slice2 := []int{}
	for {
		fmt.Print("Enter elements of slice2: ")
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}
		slice2 = append(slice2, n)
	}
	fmt.Println("slice1 :", slice1)
	fmt.Println("slice2 :", slice2)

	common := []int{}
	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				common = append(common, slice1[i])
			}
		}

	}

	fmt.Println("common :", common)
}
