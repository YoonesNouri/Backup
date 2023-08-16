package main

import "fmt"

func main() {
	var month, day, sumday int
	for {
		fmt.Print("Enter month: ")
		if _, err := fmt.Scanln(&month); err != nil || month < 1 || month > 12 {
			fmt.Println("Invalid input: (0 < month < 13)")
			continue
		}

		break
	}

	daysInMonth := []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

	for {
		fmt.Print("Enter day: ")
		if _, err := fmt.Scanln(&day); err != nil || day < 1 || day > daysInMonth[month-1] {
			fmt.Printf("Invalid input: (0 < day < %d)\n", daysInMonth[month-1]+1)
			continue
		}

		break
	}

	//هسته ی محاسباتی کد اینجاست که خیلی ساده و کوتاه شده
	//بالا دنگ و فنگهای دیتای صحیح گرفتن از یوزر است
	sumday = day
	for i := 0; i < month-1; i++ {
		sumday += daysInMonth[i]
	}

	fmt.Printf("sumday = %v\n", sumday)
}

//package main

//import "fmt"

//var d, m, numday int

//func main() {

//	fmt.Print("enter month number (1...12):")
//	for {
//		fmt.Scanln(&m)
//		if m >= 1 && m <= 12 {
//			break
//		}
//		fmt.Println("month should be between 1 and 12:")
//		fmt.Print("enter month number (1...12):")
//	}

//	if m == 12 {
//		fmt.Print("enter day number (1...29):")
//		for {
//			fmt.Scanln(&d)
//			if d >= 1 && d <= 29 {
//				break
//			}
//			fmt.Println("day should be between 1 and 29 when month is 12:")
//			fmt.Print("enter day number (1...29):")
//		}
//		numday = 6*31 + 5*30 + d
//	} else if m <= 6 {
//		fmt.Print("enter day number (1...31):")
//		for {
//			fmt.Scanln(&d)
//			if d >= 1 && d <= 31 {
//				break
//			}
//			fmt.Println("dday should be between 1 and 31 when month <= 6:")
//			fmt.Print("enter day number (1...31):")
//		}
//		numday = 31*(m-1) + d
//	} else if m >= 7 {
//		fmt.Print("enter day number (1...30):")
//		for {
//			fmt.Scanln(&d)
//			if d >= 1 && d <= 30 {
//				break
//			}
//			fmt.Println("day should be between 1 and 30 when month >= 7:")
//			fmt.Print("enter day number (1...30):")
//		}
//		numday = 6*31 + 30*(m-7) + d
//	}
//	fmt.Println(numday)
//}

//package main

//import "fmt"

//var d, m, numday int

//func main() {
//	fmt.Print("Enter month number (1...12): ")
//	fmt.Scanln(&m)

//	for {
//		fmt.Print("Enter day number: ")
//		fmt.Scanln(&d)

//		if m == 12 && (d < 1 || d > 29) {
//			fmt.Println("Invalid day. Day should be between 1 and 29 when month = 12.")
//		} else if m <= 6 && (d < 1 || d > 31) {
//			fmt.Println("Invalid day. Day should be between 1 and 31 when month <= 6.")
//		} else if m >= 7 && (d < 1 || d > 30) {
//			fmt.Println("Invalid day. Day should be between 1 and 30 when month >= 7.")
//		} else {
//			break
//		}
//	}

//	if m == 12 {
//		numday = 6*31 + 5*30 + d
//	} else if m <= 6 {
//		numday = 31*(m-1) + d
//	} else if m >= 7 {
//		numday = 6*31 + 30*(m-7) + d
//	}

//	fmt.Println(numday)
//}

// with function and map
//package main

//import "fmt"

//func main() {
//	var month, day int
//	fmt.Print("Enter the month (1-12): ")
//	for {
//		fmt.Scanln(&month)
//		if month >= 1 && month <= 12 {
//			break
//		}
//		fmt.Println("Invalid month! Please enter a value between 1 and 12.")
//		fmt.Print("Enter the month (1-12): ")
//	}

//	fmt.Print("Enter the day (1-31): ")
//	for {
//		fmt.Scanln(&day)
//		if day >= 1 && day <= getDaysInMonth(month) {
//			break
//		}
//		fmt.Printf("Invalid day! Please enter a value between 1 and %d.\n", getDaysInMonth(month))
//		fmt.Print("Enter the day (1-31): ")
//	}

// Calculate the number of days since the beginning of the year
//	numDays := day
//	for i := 1; i < month; i++ {
//		numDays += getDaysInMonth(i)
//	}

//	fmt.Printf("Number of days since the beginning of the year: %d\n", numDays)
//}

//func getDaysInMonth(month int) int {
//	daysInMonth := map[int]int{
//		1:  31,
//		2:  31,
//		3:  31,
//		4:  31,
//		5:  31,
//		6:  31,
//		7:  30,
//		8:  30,
//		9:  30,
//		10: 30,
//		11: 30,
//		12: 29,
//	}
//	return daysInMonth[month]
//}
