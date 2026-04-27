package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(s.subarraySum([]int{10, 2, -2, -20, 10}, -10))
	fmt.Println(s.subarraySum([]int{5, 1, 2, -3, 4, -2}, 3))
}

type Solution struct{}

func (s *Solution) subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixSumMap := map[int]int{0: 1}
	
	for _, num := range nums {
		sum += num
		
		if val, exists := prefixSumMap[sum-k]; exists {
			count += val
		}
		
		prefixSumMap[sum]++
	}
	
	return count
}
