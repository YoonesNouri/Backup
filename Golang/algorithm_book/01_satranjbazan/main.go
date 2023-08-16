//اگر در مسابقات شطرنج سی و دو بازیکن شرکت کنند
//و مسابقات بصورت تک حذفی برگزار شود
//تا مشخص شدن قهرمان، چند بازی باید انجام شود؟

package main

import "fmt"

func main() {
	players := 32
	games := 0

	//while
	for players > 1 {
		matches := players / 2
		games += matches
		players = matches
	}

	fmt.Println("Total number of games until the champion is determined:", games)
}
