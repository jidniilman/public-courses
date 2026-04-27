package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.reverseString("Hello, World!"))
	fmt.Println(s.reverseString("OpenAI"))
	fmt.Println(s.reverseString("Stacks are fun!"))
}

type Solution struct{}

func (s *Solution) reverseString(str string) string {
	stack := []rune{}

	for _, char := range str {
		stack = append(stack, char)
	}

	result := []rune{}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return string(result)
}
