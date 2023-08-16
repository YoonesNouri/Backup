// عدد تکراری و دفعات تکرارش در اسلایس را در خروجی چاپ کن
package main

import (
	"fmt"
)

var a int

func main() {
	A := []int{}
	repNum := []int{}
	frequency := []int{}

	for {
		fmt.Print("Enter a number : ")
		_, err := fmt.Scanln(&a)
		if err != nil {
			break
		}
		A = append(A, a)
	}
	fmt.Println("A:", A)

	for i := 0; i < len(A); i++ {
		if A[i] != 0 {
			count := 1
			for j := i + 1; j < len(A); j++ {
				if A[i] == A[j] {
					count++
					A[j] = 0 //برای شمرده نشدن مجدد
					//چون لوپ فقط برای مقادیری که مساوی با صفر نیستند هست
					//که با یک ایف این شرط را اعمال کرده
				}
			}
			if count > 1 { //کاونت یک  باشد یعنی تکرار نداریم و اساسا بیشتر از یک معنی تکرار میدهد
				repNum = append(repNum, A[i])
				frequency = append(frequency, count)
			}
		}
	}
	for i := 0; i < len(repNum); i++ {
		fmt.Printf("%v repeated %v times\n", repNum[i], frequency[i])
	}

}
