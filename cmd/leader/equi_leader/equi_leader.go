package main

import "fmt"

func equiLeader(A []int) int {
	if len(A) == 1 {
		return 0
	}
	size, value := 0, -1
	for i := range A {
		if size == 0 {
			size++
			value = A[i]
		} else {
			if value != A[i] {
				size--
			} else {
				size++
			}
		}
	}
	var candidate int
	if size > 0 {
		candidate = value
	} else {
		return 0
	}
	var count int
	for _, e := range A {
		if e == candidate {
			count++
		}
	}
	if count <= len(A)/2 {
		return 0
	}
	leader := candidate
	prefix_sums := make([]int, len(A)+1)
	for i := range A {
		if A[i] == leader {
			prefix_sums[i+1] = prefix_sums[i] + 1
		} else {
			prefix_sums[i+1] = prefix_sums[i]
		}
	}
	leader_count := 0
	for i := range prefix_sums {
		if prefix_sums[i] > i/2 && prefix_sums[len(prefix_sums)-1]-prefix_sums[i] > (len(prefix_sums)-(i+1))/2 {
			leader_count++
		}
	}
	return leader_count
}

func main() {

	A := []int{4, 3, 4, 4, 4, 2}
	fmt.Println(equiLeader(A))
}
