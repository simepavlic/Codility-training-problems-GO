package main

import "fmt"

func brackets(S string) int {
	opened := make([]rune, 0)
	brackets := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	for _, e := range S {
		if e == '(' || e == '[' || e == '{' {
			opened = append(opened, e)
		} else {
			if len(opened) == 0 || e != brackets[opened[len(opened)-1]] {
				return 0
			}
			opened = opened[:len(opened)-1]
		}
	}
	if len(opened) > 0 {
		return 0
	}
	return 1
}

func main() {

	S := "{[()()]}"
	fmt.Println(brackets(S))
	S2 := "([)()]"
	fmt.Println(brackets(S2))
}
