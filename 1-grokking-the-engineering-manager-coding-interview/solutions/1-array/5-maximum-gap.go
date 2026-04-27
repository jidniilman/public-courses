package main

import "fmt"

func main() {
	fmt.Println(findmaximumGap([]int{300, 10, 200, 100, 600}))
}

func findmaximumGap(nums []int) int {
	// Sort nums using insertion sort
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	// Find maximum gap between consecutive elements
	maxGap := 0
	for i := 1; i < len(nums); i++ {
		gap := nums[i] - nums[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}

	return maxGap
}
