package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{3, 1, 4, 2, 2}))
}

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = nums[i]
		if i > 0 {
			result[i] += result[i-1]
		}
	}
	return result
}
