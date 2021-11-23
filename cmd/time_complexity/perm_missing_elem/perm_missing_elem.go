package main

import "fmt"

func missing_elem(A []int) int {
	n := len(A) + 1
	result := n * (n + 1) / 2
	for _, e := range A {
		result -= e
	}
	return result
}

func main() {

	A := []int{2, 3, 1, 5}
	fmt.Println(missing_elem(A))
}
