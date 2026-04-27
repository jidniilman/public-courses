package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.isValid("{[()]}"))
	fmt.Println(s.isValid("{[}]"))
	fmt.Println(s.isValid("(]"))
}

type Solution struct{}

func (s *Solution) isValid(str string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
