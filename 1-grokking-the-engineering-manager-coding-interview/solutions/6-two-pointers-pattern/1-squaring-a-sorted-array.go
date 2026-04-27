package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.sortedSquares([]int{-2, -1, 0, 2, 3}))
	fmt.Println(s.sortedSquares([]int{-3, -1, 0, 1, 2}))
}

type Solution struct{}

func (s *Solution) sortedSquares(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	left := 0
	right := n - 1
	highestIndex := n - 1

	for left <= right {
		leftSquare := arr[left] * arr[left]
		rightSquare := arr[right] * arr[right]

		if leftSquare > rightSquare {
			result[highestIndex] = leftSquare
			left++
		} else {
			result[highestIndex] = rightSquare
			right--
		}
		highestIndex--
	}

	return result
}
