package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.removeDuplicateLetters("babac"))
	fmt.Println(s.removeDuplicateLetters("zabccde"))
	fmt.Println(s.removeDuplicateLetters("mnopmn"))
}

type Solution struct{}

func (s *Solution) removeDuplicateLetters(str string) string {
	lastOccurrence := make(map[rune]int)
	for i, ch := range str {
		lastOccurrence[ch] = i
	}

	stack := []rune{}
	inStack := make(map[rune]bool)

	for i, ch := range str {
		if inStack[ch] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && lastOccurrence[stack[len(stack)-1]] > i {
			removed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inStack[removed] = false
		}

		stack = append(stack, ch)
		inStack[ch] = true
	}

	return string(stack)
}
