package main

import "fmt"

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func getBiggerNumber[T numbers](t1, t2 T) T {
	if t1 == t2 {
		return t1
	}
	return t2
}

func getBiggerNumberWithComparable[T comparable](t1, t2 T) T {
	if t1 == t2 { // ./generics-sample.go:17:5: invalid operation: cannot compare t1 > t2 (operator > not defined on T)
		return t1
	}
	return t2
}

func main() {
	fmt.Println(getBiggerNumber(2.5, -4.0))
	fmt.Println(getBiggerNumberWithComparable(2.5, -4.0))
}
