package main

import "fmt"

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 1, 1, 1, 2}))
}

func minIncrementForUnique(nums []int) int {
	// Sort the array first to make it easier to ensure uniqueness
	nums = sort(nums)
	// Ensure uniqueness by incrementing duplicates and count the moves
	moves := 0
	for {
		isUnique := true
		for i := 0; i < len(nums); i++ {
			if (i != 0) && (nums[i] == nums[i-1]) {
				nums[i]++
				moves++
				isUnique = false
			}
		}
		if isUnique {
			return moves
		} else {
			// Sort and restart
			nums = sort(nums)
		}
	}
}

func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
