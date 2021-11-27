package main

import "fmt"

func min_avg(A []int) int {
	prefix_sums := make([]int, len(A)+1)

	for i := range A {
		prefix_sums[i+1] = prefix_sums[i] + A[i]
	}

	// slice with min average is either 2 or 3 elements long
	// because slice of every other length can be mad out of these two
	min_average, min_index := 10000., 0
	var two_avg, three_avg float64
	two_avg = float64(prefix_sums[2]-prefix_sums[0]) / 2
	if two_avg < min_average {
		min_average = two_avg
		min_index = 0
	}
	for i := 0; i < len(prefix_sums)-3; i++ {
		two_avg = float64(prefix_sums[i+3]-prefix_sums[i+1]) / 2
		three_avg = float64(prefix_sums[i+3]-prefix_sums[i]) / 3
		if two_avg < min_average {
			min_average = two_avg
			min_index = i + 1
		}
		if three_avg < min_average {
			min_average = three_avg
			min_index = i
		}
	}

	return min_index
}

func main() {

	A := []int{4, 2, 2, 5, 1, 5, 8}
	fmt.Println(min_avg(A))
}
