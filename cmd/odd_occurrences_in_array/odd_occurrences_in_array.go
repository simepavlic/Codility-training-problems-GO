package main

import (
	"errors"
	"fmt"
	"os"
)

func unpaired_element(A []int) (int, error) {
	if len(A)%2 == 0 {
		return -1, errors.New("Number of elements must be odd")
	}
	elem_count := make(map[int]int, 0)
	for _, e := range A {
		elem_count[e]++
	}
	var odd_elem int
	for k, v := range elem_count {
		if v%2 == 1 {
			odd_elem = k
			return odd_elem, nil
		}
	}
	return -1, nil
}

func xor_unpaired_element(A []int) (int, error) {
	if len(A)%2 == 0 {
		return -1, errors.New("Number of elements must be odd")
	}
	var result int
	for _, e := range A {
		result ^= e
	}
	return result, nil
}

func main() {

	A := []int{9, 3, 9, 3, 9, 7, 9, 7, 2}
	e, err := xor_unpaired_element(A)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(e)
}
