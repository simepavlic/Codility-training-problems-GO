package main

import "fmt"

func fish(A []int, B []int) int {
	fish_stack := make([]int, 0)
	var upstream_alive int
	for i := range B {
		switch B[i] {
		case 0:
			for len(fish_stack) > 0 {
				if A[i] > fish_stack[len(fish_stack)-1] {
					fish_stack = fish_stack[:len(fish_stack)-1]
				} else {
					break
				}
			}
			if len(fish_stack) == 0 {
				upstream_alive++
			}
		case 1:
			fish_stack = append(fish_stack, A[i])
		}
	}
	return upstream_alive + len(fish_stack)
}

func main() {

	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}
	fmt.Println(fish(A, B))
}
