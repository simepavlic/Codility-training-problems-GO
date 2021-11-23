package main

import "fmt"

func distinct(A []int) int {
	seen := make(map[int]bool)
	var distinct int
	for _, e := range A {
		if _, ok := seen[e]; !ok {
			distinct++
			seen[e] = true
		}
	}
	return distinct
}

func main() {

	A := []int{2, 1, 1, 2, 3, 1}
	fmt.Println(distinct(A))
}
