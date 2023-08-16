// کتاب مسائل زیادی از دنباله ها آورده که ارزش کد زدن نداره چون نکته ی خاصی نداره
// i/(Fac(n)) > 0.71 این رو بعنوان نمونه حل میکنم، مجموع چند جمله بزرگتر از این عدد میشود
package main

import (
	"fmt"
)

func Fac(n int) int {
	fc := 1
	for i := 1; i <= n; i++ {
		fc *= i
	}
	return fc
}

func main() {

	sum := 0.0
	i := 2
	for {
		sum += 1.0 / float64(Fac(i))
		if sum > 0.71 {
			fmt.Printf("sum of %d-th sentence > 0.71", i-1)
			return // exits form func main()
		}
		i++
	}
}
