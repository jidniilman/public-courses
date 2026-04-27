package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.countElements([]int{4, 3, 1, 5, 6}))
	fmt.Println(s.countElements([]int{7, 8, 9, 10}))
	fmt.Println(s.countElements([]int{11, 13, 15, 16}))
}

type Solution struct{}

func (s *Solution) countElements(arr []int) int {
	numSet := make(map[int]bool)
	for _, num := range arr {
		numSet[num] = true
	}

	count := 0
	for _, num := range arr {
		if numSet[num+1] {
			count++
		}
	}

	return count
}
