package main

import (
	"fmt"
	"strconv"
)

func binaryGap(N int) int {
	binary_n := strconv.FormatInt(int64(N), 2)
	longest_gap, current_gap := 0, 0
	for _, char := range binary_n {
		switch char {
		case '0':
			current_gap++
		case '1':
			if current_gap > 0 {
				if current_gap > longest_gap {
					longest_gap = current_gap
				}
				current_gap = 0
			}
		}
	}
	return longest_gap
}

func betterBinaryGap(N int) int {
	for N%2 == 0 {
		N >>= 1
	}
	var maxSteps, steps int
	for ; N != 0; N >>= 1 {
		switch N % 2 {
		case 0:
			steps++
		case 1:
			if steps > 0 {
				if steps > maxSteps {
					maxSteps = steps
				}
				steps = 0
			}
		}
	}
	return maxSteps
}

func bestBinaryGap(N int) int {
	var steps int
	for N%2 == 0 {
		N >>= 1
	}
	for (N & (N + 1)) != 0 {
		N |= N >> 1
		steps++
	}
	return steps
}

func main() {

	number := 9
	fmt.Println(binaryGap(number))
	fmt.Println(betterBinaryGap(number))
	fmt.Println(bestBinaryGap(number))

}
