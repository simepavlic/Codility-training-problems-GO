package main

import "fmt"

func max_counters(N int, A []int) []int {
	counters := make([]int, N)
	var maxCounter, lastUpdate int
	for _, e := range A {
		if e == N+1 {
			lastUpdate = maxCounter
		} else {
			if counters[e-1] < lastUpdate {
				counters[e-1] = lastUpdate + 1
			} else {
				counters[e-1]++
			}
			if counters[e-1] > maxCounter {
				maxCounter = counters[e-1]
			}
		}
	}
	for i := range counters {
		if counters[i] < lastUpdate {
			counters[i] = lastUpdate
		}
	}
	return counters
}

func main() {

	A := []int{3, 4, 4, 6, 1, 4, 4}
	fmt.Println(max_counters(5, A))
}
