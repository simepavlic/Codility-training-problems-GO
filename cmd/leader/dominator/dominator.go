package main

import "fmt"

func dominator(A []int) int {
	if len(A) == 0 {
		return -1
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
		return -1
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
