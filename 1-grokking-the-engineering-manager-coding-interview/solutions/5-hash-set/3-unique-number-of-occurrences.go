package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.uniqueOccurrences([]int{4, 5, 4, 6, 6, 6}))
	fmt.Println(s.uniqueOccurrences([]int{7, 8, 8, 9, 9, 9, 10, 10}))
	fmt.Println(s.uniqueOccurrences([]int{11, 12, 13, 14, 14, 15, 15, 15}))
}

type Solution struct{}

func (s *Solution) uniqueOccurrences(arr []int) bool {
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	occurrenceSet := make(map[int]bool)
	for _, count := range countMap {
		if occurrenceSet[count] {
			return false
		}
		occurrenceSet[count] = true
	}

	return true
}
