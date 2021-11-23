package main

import "fmt"

func impact_factor(r byte) int {
	switch r {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	}
	return 0
}

func genomic_range(S string, P []int, Q []int) []int {
	results := make([]int, len(P))
	var left, right int
	for i := range P {
		left, right = P[i], Q[i]
		min_impact := 4
		for j := left; j <= right; j++ {
			impact := impact_factor(S[j])
			if impact < min_impact {
				min_impact = impact
				if min_impact == 1 {
					break
				}
			}
		}
		results[i] = min_impact
	}
	return results
}

func genomic_range_prefix_sums(S string, P []int, Q []int) []int {
	sums := make([][]int, 3)
	for i := 0; i < 3; i++ {
		sums[i] = make([]int, len(S)+1)
	}
	for i, e := range S {
		var a, c, g int
		switch e {
		case 'A':
			a = 1
		case 'C':
			c = 1
		case 'G':
			g = 1
		}
		sums[0][i+1] = sums[0][i] + a
		sums[1][i+1] = sums[1][i] + c
		sums[2][i+1] = sums[2][i] + g
	}
	result := make([]int, len(P))
	for i := range P {
		left, right := P[i], Q[i]+1
		if sums[0][right]-sums[0][left] > 0 {
			result[i] = 1
		} else if sums[1][right]-sums[1][left] > 0 {
			result[i] = 2
		} else if sums[2][right]-sums[2][left] > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}
	return result
}

func main() {

	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}

	fmt.Println(genomic_range_prefix_sums(S, P, Q))
}
