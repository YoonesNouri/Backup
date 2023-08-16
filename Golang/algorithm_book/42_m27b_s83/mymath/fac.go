package mymath

func Fac(n int) int {
	fc := 1
	for i := 1; i <= n; i++ {
		fc *= i
	}
	return fc
}
