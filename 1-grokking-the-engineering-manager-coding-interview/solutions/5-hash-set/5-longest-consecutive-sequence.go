package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.longestConsecutive([]int{10, 11, 14, 12, 13}))
	fmt.Println(s.longestConsecutive([]int{3, 6, 4, 100, 101, 102}))
	fmt.Println(s.longestConsecutive([]int{4, 3, 6, 2, 5, 8, 4, 7, 0, 1}))
	fmt.Println(s.longestConsecutive([]int{7, 8, 10, 11, 15}))
}

type Solution struct{}

func (s *Solution) longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}
