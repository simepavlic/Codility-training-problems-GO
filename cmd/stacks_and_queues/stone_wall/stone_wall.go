package main

import (
	"fmt"
	"sort"
)

func stone_wall(H []int) int {
	sort.Ints(H)
	var count int
	for i := range H[:len(H)-1] {
		if H[i+1] > H[1] {
			count++
		}
	}
	return count
}

func main() {
	H := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}
	fmt.Println(stone_wall(H))
}
