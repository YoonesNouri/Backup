package main

import "fmt"

func main() {
	var dayNumber int
	fmt.Print("Enter the day number after the first day of the year: ")
	fmt.Scanln(&dayNumber)

	// Define the number of days in each month
	daysInMonth := []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

	// Validate the day number
	if dayNumber < 1 || dayNumber > 365 {
		fmt.Println("Invalid day number.")
		return
	}

	monthname := map[int]string{1: "فروردین", 2: "اردیبهشت", 3: "خرداد", 4: "تیر", 5: "مرداد", 6: "شهریور", 7: "مهر", 8: "آبان", 9: "آذر", 10: "دی", 11: "بهمن", 12: "اسفند"}

	// Find the corresponding month and day
	month := 0
	for i, days := range daysInMonth {
		if dayNumber <= days {
			month = i + 1
			break
		}
		dayNumber -= days
	}

	// Print the corresponding day and month
	fmt.Printf("%v م%v", monthname[month], dayNumber)
}
