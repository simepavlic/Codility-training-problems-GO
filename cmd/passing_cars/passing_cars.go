package main

import "fmt"

func passing_cars(A []int) int {
	var counter int
	for i, e := range A {
		if e == 0 {
			for _, e2 := range A[i:] {
				if e2 == 1 {
					counter++
					if counter > 1000000000 {
						return -1
					}
				}
			}
		}
	}
	return counter
}

func passing_cars_better(A []int) int {
	var counter, num_west int
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 1 {
			num_west++
		}
		if A[i] == 0 {
			counter += num_west
			if counter > 1000000000 {
				return -1
			}
		}
	}
	return counter
}

func main() {
	A := []int{0, 1, 0, 1, 1}
	fmt.Println(passing_cars_better(A))
}
