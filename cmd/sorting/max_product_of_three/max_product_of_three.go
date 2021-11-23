package main

import (
	"fmt"
	"sort"
)

func max_product_of_three(A []int) int {
	sort.Ints(A)
	l := len(A)
	// either last 3 or first 2 + last (when we have big negatives)
	first := A[l-1] * A[l-2] * A[l-3]
	second := A[l-1] * A[0] * A[1]
	if first > second {
		return first
	}
	return second
}

func main() {

	A := []int{-3, 1, 2, -2, 5, 6}
	fmt.Println(max_product_of_three(A))
}
