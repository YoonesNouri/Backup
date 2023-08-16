package main

import (
	"fmt"
)

func main() {
	var num int
	var numbers []int

	for {
		fmt.Print("Enter an integer (press enter to exit): ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
	}

	// Create a map to find the frequency of each element
	frequency := make(map[int]bool) //وقتی مپ خالی است پس ولیوهای بول صفر=فالس است
	uniquenumbers := make([]int, 0)

	// Iterate over the numbers slice and check for repeated elements
	for _, n := range numbers {
		if !frequency[n] { //کی/اِن رو میده ولیو/بول رو چک میکنه ، اگه اِن در مپ نبود
						   //بولش میشه فالس و ضدش میشه ترو
			frequency[n] = true                      // بول اِن رو تغییر میده به ترو
			uniquenumbers = append(uniquenumbers, n) //اِن رو اپند/تعبیه میکنه تو یه اسلایس
		} //اگه بازم همون اِن قبلی اومد بولش ترو هست و ضدش فالس هست پس اپند نمیشه = تکراری حذف میشه
	}

	fmt.Println("Remaining Elements:")
	fmt.Println(uniquenumbers)
}
