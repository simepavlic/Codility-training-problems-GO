package main

import "fmt"

func maxSlice(A []int) int {
	maxCurrent, maxSlice := -2147483648, -2147483648
	for _, e := range A {
		current := maxCurrent + e
		if current > e {
			maxCurrent = current
		} else {
			maxCurrent = e
		}
		if maxSlice < maxCurrent {
			maxSlice = maxCurrent
		}
	}
	return maxSlice
}

func main() {

	A := []int{3, 2, -6, 4, 0}
	fmt.Println(maxSlice(A))
}
