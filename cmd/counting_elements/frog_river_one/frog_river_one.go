package main

import "fmt"

func frog_river_one(X int, A []int) int {
	position_map := make(map[int]int, X)
	for i := 1; i <= X; i++ {
		position_map[i] = 1
	}

	for i := 0; i < len(A); i++ {
		_, ok := position_map[A[i]]
		if ok {
			delete(position_map, A[i])
			if len(position_map) == 0 {
				return i
			}
		}
	}
	return -1
}

func frog_river_one_steps(X int, A []int) int {
	positions := make([]bool, X)
	steps := X

	for i := 0; i < len(A); i++ {
		if !positions[A[i]-1] {
			positions[A[i]-1] = true
			steps--
			if steps == 0 {
				return i
			}
		}
	}
	return -1
}

func main() {

	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	fmt.Println(frog_river_one_steps(5, A))

}
