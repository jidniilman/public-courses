package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.decimalToBinary(2))
	fmt.Println(s.decimalToBinary(7))
	fmt.Println(s.decimalToBinary(18))
}

type Solution struct{}

func (s *Solution) decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	stack := []int{}

	for n > 0 {
		remainder := n % 2
		stack = append(stack, remainder)
		n = n / 2
	}

	result := ""
	for len(stack) > 0 {
		if stack[len(stack)-1] == 0 {
			result += "0"
		} else {
			result += "1"
		}
		stack = stack[:len(stack)-1]
	}

	return result
}
