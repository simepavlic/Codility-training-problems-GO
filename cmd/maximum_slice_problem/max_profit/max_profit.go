package main

import "fmt"

func maxProfit(A []int) int {
	min, maxProfit := 400000, 0
	for _, e := range A {
		if min > e {
			min = e
		}
		profit := e - min
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func main() {

	A := []int{23171, 21011, 21123, 21366, 21013, 21367}
	fmt.Println(maxProfit(A))
}
