package main

import "fmt"

func count_div(A int, B int, K int) int {
	var result int
	if A%K == 0 {
		result++
	}
	return result + B/K - A/K
}

func main() {

	A, B, K := 6, 11, 2
	fmt.Println(count_div(A, B, K))
}
