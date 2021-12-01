package main

import "fmt"

func maxDoubleSliceSum(A []int) int {
	// TODO analyze this for errors, try to fix
	min_between := 10000
	maxCurrent, maxFull, maxSlice := 0, 0, 0
	for i := 1; i < len(A)-1; i++ {
		if min_between > A[i] {
			min_between = A[i]
		}
		maxFull += A[i]
		if maxFull-min_between >= 0 {
			maxCurrent = maxFull - min_between
		} else {
			maxCurrent = 0
			maxFull = 0
			min_between = 10000
		}
		if maxSlice < maxCurrent {
			maxSlice = maxCurrent
		}
	}
	return maxSlice
}

func maxDoubleSliceSumLeftRight(A []int) int {
	maxCurrent := 0
	n := len(A)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)
	for i := 1; i < n-1; i++ {
		current := maxCurrent + A[i]
		if current > 0 {
			maxCurrent = current
		} else {
			maxCurrent = 0
		}

		maxLeft[i] = maxCurrent
	}
	maxCurrent = 0
	for i := n - 2; i > 0; i-- {
		current := maxCurrent + A[i]
		if current > 0 {
			maxCurrent = current
		} else {
			maxCurrent = 0
		}

		maxRight[i] = maxCurrent
	}
	maxDoubleSlice, currentDoubleSlice := 0, 0
	for i := 1; i < n-1; i++ {
		currentDoubleSlice = maxLeft[i-1] + maxRight[i+1]
		if currentDoubleSlice > maxDoubleSlice {
			maxDoubleSlice = currentDoubleSlice
		}
	}
	return maxDoubleSlice
}

func main() {

	A := []int{3, 2, 6, -1, 4, 5, -1, 2}
	fmt.Println(maxDoubleSliceSumLeftRight(A))
}
