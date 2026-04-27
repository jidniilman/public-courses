package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.continuousSubarrays([]int{1, 2, 2, 13, 1}))
	fmt.Println(s.continuousSubarrays([]int{3, 3, 4, 2}))
	fmt.Println(s.continuousSubarrays([]int{5, 17, 35, 26, 5}))
}

type Solution struct{}

func Constructor() Solution {
	return Solution{}
}

func (s *Solution) continuousSubarrays(nums []int) int64 {
	count := int64(0)
	left := 0
	minDeque := []int{}
	maxDeque := []int{}

	for right := 0; right < len(nums); right++ {
		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] >= nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, right)

		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] <= nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, right)

		for nums[maxDeque[0]]-nums[minDeque[0]] > 2 {
			left++
			if minDeque[0] < left {
				minDeque = minDeque[1:]
			}
			if maxDeque[0] < left {
				maxDeque = maxDeque[1:]
			}
		}

		count += int64(right - left + 1)
	}

	return count
}
