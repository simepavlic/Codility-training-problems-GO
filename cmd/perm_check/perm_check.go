package main

import "fmt"

func is_permutation(A []int) bool {
	seen_elems := make([]bool, len(A))
	for _, e := range A {
		if e > len(A) || e < 1 || seen_elems[e-1] {
			return false
		}
		seen_elems[e-1] = true
	}
	return true
}

func main() {

	A := []int{4, 1, 3, 2, 5, 7}
	fmt.Println("Array is permutation? ", is_permutation(A))
}
