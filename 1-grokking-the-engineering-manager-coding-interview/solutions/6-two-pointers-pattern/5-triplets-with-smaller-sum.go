package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.threeSumSmaller([]int{-1, 0, 2, 3}, 3))
	fmt.Println(s.threeSumSmaller([]int{-1, 4, 2, 1, 3}, 5))
}

type Solution struct{}

func (s *Solution) sortInts(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func (s *Solution) threeSumSmaller(arr []int, target int) int {
	s.sortInts(arr)
	count := 0

	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]

			if sum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}
