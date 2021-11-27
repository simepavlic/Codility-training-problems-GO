package main

import "fmt"

func dominator(A []int) int {
	if len(A) == 0 {
		return -1
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
	var count, index int
	for i, e := range A {
		if e == candidate {
			count++
			index = i
		}
	}
	if count > len(A)/2 {
		return index
	}
	return -1
}

func main() {

	A := []int{3, 4, 3, 2, 3, -1, 3, 3}
	fmt.Println(dominator(A))
}
