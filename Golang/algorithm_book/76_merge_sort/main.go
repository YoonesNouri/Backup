package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	merged := make([]int, size)

	for k := 0; k < size; k++ {
		if i >= len(left) {
			merged[k] = right[j]
			j++
		} else if j >= len(right) {
			merged[k] = left[i]
			i++
		} else if left[i] < right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
	}

	return merged
}

func main() {
	arr := []int{9, 2, 7, 1, 25, 6, 8, 3, 2, 3, 4, 5, 7, 1, 39, 2, 6, 8, 2, 8, 4}
	sorted := mergeSort(arr)
	fmt.Println("Sorted array:", sorted)
}
