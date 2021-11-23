package main

import (
	"fmt"
)

func cyclicRotation(A []int, K int) []int {
	if len(A) == 0 || len(A) == 1 {
		return A
	}
	K %= len(A)
	return append(A[len(A)-K:], A[:len(A)-K]...)
}

func main() {

	A := []int{3, 8, 9, 7, 6}
	fmt.Println(cyclicRotation(A, 3))
}
