package digits

func RevSplitInt(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed
}
