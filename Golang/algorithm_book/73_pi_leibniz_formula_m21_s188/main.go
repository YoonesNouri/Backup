// π = 4 * (1 - 1/3 + 1/5 - 1/7 + ... + (-1)^n/(2n+1) + ...)
// عدد پی را با استفاده از فرمول لایبنیز تا 50 رقم اعشار محاسبه کنید
package main

import (
	"fmt"
)

func calculatePi(iterations int) float64 {
	pi := 0.0
	sign := 1.0 //ضریب مثبت و منفی که در جمله ی اول +1 است

	for i := 0; i < iterations; i++ {
		term := 1.0 / (2.0*float64(i) + 1.0)
		pi += sign * term
		sign *= -1.0 //ضریب را در هر جمله در -1 ضرب میکنیم تا یکی درمیان مثبت و منفی شود
	}

	pi *= 4.0 // pi = 4.0 * pi
	return pi
}

func main() {
	iterations := 1000000
	pi := calculatePi(iterations)
	piFormatted := fmt.Sprintf("%.50f", pi)
	fmt.Println("Approximation of pi:", piFormatted)
}
