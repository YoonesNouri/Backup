package main

import (
	"fmt"
)

var name string

func main() {
	invitedGuests := make(map[string]bool)
	for {
		fmt.Print("Add a name to the list (type 'finish' to finish): ")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Print("Invalid input.")
			continue
		}
		if name == "finish" {
			break
		}
		invitedGuests[name] = true
	}

	allChecked := 0
	for {
		var name string
		fmt.Print("Insert the name you want to check (type 'finish' to exit): ")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}

		if name == "finish" {
			fmt.Println("Goodbye!")
			return //ریترن در کل برای بستن فانکشن و خروج از آن است حتی اگر در فانک مِین باشد
			//از آن هم خارج میشود و کد پایان می یابد
			//گاهی در کد نویسی چنین حرکتی ضروری است. مثلا به هر دلیلی کاربر میخواهد کد پایان یابد
			//بعد از وارد کردن مثلا فینیش از برنامه بوسیله ی ریترنی که داخل فانک مین است، کلا خارج میشود
			//اگر بریک بود اسامی مهمان ها را هم چاپ میکرد اما اینطوری چاپ نمیکند چون از کد کلا خارج میشود
		}

		if invited, ok := invitedGuests[name]; ok && invited {
			fmt.Printf("%s is invited\n", name)
			allChecked++
			if allChecked == len(invitedGuests) {
				break
			}
		} else {
			fmt.Printf("%s is not invited\n", name)
			continue
		}
	}

	fmt.Println("Invited names:")
	for guest := range invitedGuests {
		fmt.Println(guest)
	}
}
