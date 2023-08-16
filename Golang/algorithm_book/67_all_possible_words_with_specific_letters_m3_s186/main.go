package main

import "fmt"

func main() {
	letters := []rune{'I', 'R', 'A', 'N'}
	word := make([]rune, len(letters))
	used := make([]bool, len(letters)) // empty bool = false
	count := 0

	var generateWords func(int)
	generateWords = func(index int) {
		if index == len(word) { // وقتی تعداد حروف تعبیه شده داخل اسلایس ورد به تعداد دلخواه رسید یعنی کلمه تشکیل شده
			fmt.Println(string(word)) //پس رون های کنار هم تعبیه شده را بصورت تایپ استرینگ پرینت میکند تا کلمه پرینت شود
			count++                   // یک کلمه تولید شد پس یک عدد به کانت اضافه میکند
			return
		}

		for i := 0; i < len(letters); i++ {
			if !used[i] {
				word[index] = letters[i] //حروف استفاده نشده را بصورت رون در داخل اسلایس ورد تعبیه میکند
				used[i] = true           //کلمه ی استفاده شده را ترو میکند تا در دور بعد کال شدن فانکشن استفاده نشود
				generateWords(index + 1) //به ایندکس فانکشن یکی اضافه میکند تا حرف بعدی کلمه را بررسی کند
				used[i] = false          // یوزد را فالس کرد تا این حرف در ترکیبات بعدی در پوزیشنهای مختلف دیگر استفاده شود
			}
		}
	}

	generateWords(0) //ایندکس 0 را بعنوان شروع لوپ تولید کلمه میگیرد
	//تا هر چهار حرف هر کلمه را دربربگیرد

	fmt.Println("Number of possible words:", count)
}
