package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.minSubArrayLen(7, []int{2, 1, 5, 2, 3, 2}))
	fmt.Println(s.minSubArrayLen(7, []int{2, 1, 5, 2, 8}))
	fmt.Println(s.minSubArrayLen(8, []int{3, 4, 1, 1, 6}))
}

type Solution struct{}

func (s *Solution) minSubArrayLen(target int, arr []int) int {
	minLength := len(arr) + 1
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]

		for windowSum >= target {
			currentLength := windowEnd - windowStart + 1
			if currentLength < minLength {
				minLength = currentLength
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	if minLength == len(arr)+1 {
		return 0
	}
	return minLength
}
