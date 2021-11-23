package main

import "fmt"

func nesting(S string) int {
	var opened int
	for _, e := range S {
		switch e {
		case '(':
			opened++
		case ')':
			if opened == 0 {
				return 0
			}
			opened--
		}
	}
	if opened > 0 {
		return 0
	}
	return 1
}

func main() {
	S := "(()(())())"
	fmt.Println(nesting(S))
	S2 := "())"
	fmt.Println(nesting(S2))
}
