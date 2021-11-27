package main

import "fmt"

func equiLeader(A []int) int {
	if len(A) == 1 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, A[0])
	for i := 1; i < len(A); i++ {
		stack = append(stack, A[i])
		if len(stack) > 1 {
			if stack[len(stack)-1] != stack[len(stack)-2] {
				stack = stack[:len(stack)-2]
			}
		}
	}
	var candidate int
	if len(stack) > 0 {
		candidate = stack[0]
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
