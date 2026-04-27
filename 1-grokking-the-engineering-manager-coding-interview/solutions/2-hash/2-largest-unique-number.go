package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.largestUniqueNumber([]int{5, 7, 3, 7, 5, 8}))
	fmt.Println(s.largestUniqueNumber([]int{1, 2, 3, 2, 1, 4, 4}))
	fmt.Println(s.largestUniqueNumber([]int{9, 9, 8, 8, 7, 7}))
}

type Solution struct{}

func (s *Solution) largestUniqueNumber(A []int) int {
	frequencyMap := map[int]int{}

	for _, num := range A {
		frequencyMap[num]++
	}

	maxUnique := -1
	for num, count := range frequencyMap {
		if count == 1 && num > maxUnique {
			maxUnique = num
		}
	}

	return maxUnique
}
