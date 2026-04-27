package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.findLeastNumOfUniqueInts([]int{5, 5, 4, 3, 2, 3, 2, 3, 3, 2}, 6))
	fmt.Println(s.findLeastNumOfUniqueInts([]int{7, 7, 7, 8, 8, 9}, 2))
	fmt.Println(s.findLeastNumOfUniqueInts([]int{1, 2, 2, 3, 4, 3}, 4))
}

type Solution struct{}

func (s *Solution) findLeastNumOfUniqueInts(arr []int, k int) int {
	frequencyMap := map[int]int{}

	for _, num := range arr {
		frequencyMap[num]++
	}

	frequencies := []int{}
	for _, count := range frequencyMap {
		frequencies = append(frequencies, count)
	}

	for i := 0; i < len(frequencies); i++ {
		for j := i + 1; j < len(frequencies); j++ {
			if frequencies[i] > frequencies[j] {
				frequencies[i], frequencies[j] = frequencies[j], frequencies[i]
			}
		}
	}

	removals := k
	uniqueCount := len(frequencies)

	for _, freq := range frequencies {
		if removals >= freq {
			removals -= freq
			uniqueCount--
		} else {
			break
		}
	}

	return uniqueCount
}
