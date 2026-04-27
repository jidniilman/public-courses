package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.dailyTemperatures([]int{70, 73, 75, 71, 69, 72, 76, 73}))
	fmt.Println(s.dailyTemperatures([]int{73, 72, 71, 70}))
	fmt.Println(s.dailyTemperatures([]int{70, 71, 72, 73}))
}

type Solution struct{}

func (s *Solution) dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return result
}
