package main

import (
	"fmt"
)

func frog_jmp(X int, Y int, D int) int {
	return (Y - X + D - 1) / D
}

func main() {
	fmt.Println(frog_jmp(10, 85, 30))
}
