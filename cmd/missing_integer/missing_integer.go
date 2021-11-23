package main

import (
	"fmt"
)

func missing_integer(A []int) int {
	// pigeonhole principle
	seen_integers := make([]bool, len(A)+1)
	for _, e := range A {
		if e > 0 && e < len(A)+2 && !seen_integers[e-1] {
			seen_integers[e-1] = true
		}
	}
	smallest := 1
	for i, s := range seen_integers {
		if !s {
			smallest = i + 1
			break
		}
	}
	return smallest
}

func main() {

	A := []int{1, 3, 6, 4, 1, 2}
	fmt.Println(missing_integer(A))
}
