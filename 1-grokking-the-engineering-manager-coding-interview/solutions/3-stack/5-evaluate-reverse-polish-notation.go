package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.evalRPN([]string{"2", "11", "5", "/", "+"}))
	fmt.Println(s.evalRPN([]string{"2", "3", "11", "+", "*"}))
	fmt.Println(s.evalRPN([]string{"5", "1", "2", "+", "4", "*", "+", "3", "-"}))
}

type Solution struct{}

func (s *Solution) stringToInt(str string) int {
	result := 0
	negative := false
	start := 0

	if len(str) > 0 && str[0] == '-' {
		negative = true
		start = 1
	}

	for i := start; i < len(str); i++ {
		result = result*10 + int(str[i]-'0')
	}

	if negative {
		return -result
	}
	return result
}

func (s *Solution) evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		} else {
			num := s.stringToInt(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
