package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.sumSubarrayMins([]int{3, 1, 2, 4, 5}))
	fmt.Println(s.sumSubarrayMins([]int{2, 6, 5, 4}))
	fmt.Println(s.sumSubarrayMins([]int{7, 3, 8}))
}

type Solution struct{}

func (s *Solution) sumSubarrayMins(arr []int) int {
	mod := int(1e9 + 7)
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)

	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n - i
		} else {
			right[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}

	result := 0
	for i := 0; i < n; i++ {
		result = (result + arr[i]*left[i]*right[i]) % mod
	}

	return result
}
