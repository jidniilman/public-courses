package main

import (
	"fmt"
	"sort"
)

func main() {
	s := Solution{}
	fmt.Println(s.findLongestChain([][]int{{1, 2}, {3, 4}, {2, 3}}))
	fmt.Println(s.findLongestChain([][]int{{5, 6}, {1, 2}, {8, 9}, {2, 3}}))
	fmt.Println(s.findLongestChain([][]int{{7, 8}, {5, 6}, {1, 2}, {3, 5}, {4, 5}, {2, 3}}))
}

type Solution struct{}

func (s *Solution) findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	count := 1
	currentEnd := pairs[0][1]

	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > currentEnd {
			count++
			currentEnd = pairs[i][1]
		}
	}

	return count
}
