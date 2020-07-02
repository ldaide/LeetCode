package main

import "fmt"

func minAddToMakeValid(S string) int {

	stack := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == ')' && len(stack) > 0 {
			e := stack[len(stack)-1]
			if e == '(' {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, S[i])
	}

	return len(stack)
}

func main() {
	s := "()))(("
	fmt.Println(minAddToMakeValid(s))
}
