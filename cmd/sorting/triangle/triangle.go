package main

import (
	"fmt"
	"sort"
)

func triangle(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Ints(A)
	for i := range A[:len(A)-2] {
		if A[i]+A[i+1] > A[i+2] && A[i+1]+A[i+2] > A[i] && A[i]+A[i+2] > A[i+1] {
			return 1
		}
	}
	return 0
}

func main() {

	A := []int{10, 2, 5, 1, 8, 20}
	fmt.Println(triangle(A))
	B := []int{10, 50, 5, 1}
	fmt.Println(triangle(B))
}
