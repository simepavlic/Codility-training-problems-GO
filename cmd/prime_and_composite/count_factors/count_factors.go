package main

import (
	"fmt"
	"math"
)

func countFactors(N int) int {
	count := 1
	for i := 2; i <= N; i++ {
		if N%i == 0 {
			count++
		}
	}
	return count
}

// factors come in pairs where 1 is lower than square root of N, and other is higher
func countFactorsBetter(N int) int {
	count := 0
	squareRoot := (int)(math.Ceil(math.Sqrt(float64(N))))
	for i := 1; i < squareRoot; i++ {
		if N%i == 0 {
			count += 2
		}
	}
	if squareRoot*squareRoot == N {
		count++
	}
	return count
}

func main() {
	fmt.Println(countFactors(24))
	fmt.Println(countFactorsBetter(24))
}
