package main

import "fmt"

func intersect(r1, r2, d int) bool {
	return r1+r2 >= d
}

func number_of_disc_intersections(A []int) int {
	var count int
	for i := range A {
		for j := i + 1; j < len(A); j++ {
			if intersect(A[i], A[j], j-i) {
				count++
			}
		}
	}
	if count > 10000000 {
		return -1
	}
	return count
}

func main() {

	A := []int{1, 5, 2, 1, 4, 0}
	fmt.Println(number_of_disc_intersections(A))

}
