package main

import (
	"testing"
)

func BenchmarkGeneratePrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generatePrimeNumbers(10000)
	}
}
