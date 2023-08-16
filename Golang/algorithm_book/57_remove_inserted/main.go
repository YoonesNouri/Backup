package main

import "fmt"

func main() {
	var numbers []int

	// Take ints and put them in a slice
	for {
		var num int
		fmt.Print("Enter an integer (or any non-integer to exit): ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			break //کانتینیو نیست که به اول لوپ برود ، بریک هست و از لوپ خارج میشود
		}
		numbers = append(numbers, num)
	}

	// Print the original slice
	fmt.Println("original Slice:", numbers)

	// Take ints again and remove them from the slice
	for {
		var num int
		fmt.Print("Enter an integer to remove (or any non-integer to exit): ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		// Find and remove the input number from the slice
		for i := 0; i < len(numbers); i++ {
			if numbers[i] == num {
				numbers = append(numbers[:i], numbers[i+1:]...)
				i-- //را یک واحد کم میکنیم تا المنت بعدی که i
				//به اندیس المنت حذف شده منتقل شده هم در چرخه حساب شود
			}
		}
	}

	// Print the remaining elements
	fmt.Println("Remaining Elements:", numbers)
}
