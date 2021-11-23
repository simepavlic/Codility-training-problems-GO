package main

import (
	"fmt"
	"math"
)

func tape_equilibrium(A []int) int {
	var left_sum int
	for _, e := range A {
		left_sum += e
	}
	min_splitter := math.MaxInt
	var right_sum int
	for i := len(A) - 1; i > 0; i-- {
		left_sum -= A[i]
		right_sum += A[i]
		abs_sum := int(math.Abs(float64(left_sum - right_sum)))
		if abs_sum < min_splitter {
			min_splitter = abs_sum
		}
	}
	return min_splitter
}

func main() {

	A := []int{-3, 1}
	fmt.Println(tape_equilibrium(A))
}
