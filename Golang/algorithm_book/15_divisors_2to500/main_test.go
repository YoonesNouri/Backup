package main

import "testing"

func TestMySum(t *testing.T) {
	for n := 2; n <= 10000; n++ {
		divisors := []int{}
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				divisors = append(divisors, i)
			}
		}
	}

}
