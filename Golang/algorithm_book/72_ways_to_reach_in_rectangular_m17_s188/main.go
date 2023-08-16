//فرض کن جدولی مستطیلی شبکه ای داریم به ابعاد ایکس در وای
//که تعداد نقطه ی تلاقی راه ها در هر ضلع را نشان میدهد
//میخواهیم از پایین چپ با بالا راست این جدول حرکت کنیم
//محاسبه کن چند راه وجود دارد برای رسیدن؟
//حرکات فقط به سمت بالا و راست هستند
//Ways = (x+y-2)! / ((x-1)! * (y-1)!)

package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func calculateWays(x, y int) int {
	numerator := factorial(x + y - 2)
	denominator := factorial(x-1) * factorial(y-1)
	ways := numerator / denominator
	return ways
}

func main() {
	var x, y int

	// Prompt the user to enter the dimensions of the rectangular table
	fmt.Print("Enter the number of houses in the x-direction: ")
	fmt.Scanln(&x)
	fmt.Print("Enter the number of houses in the y-direction: ")
	fmt.Scanln(&y)

	ways := calculateWays(x, y)
	fmt.Println("Number of ways to reach the top right:", ways)
}
