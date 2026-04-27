package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.threeSum([]int{-3, 0, 1, 2, -1, 1, -2}))
	fmt.Println(s.threeSum([]int{-5, 2, -1, -2, 3}))
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

func (s *Solution) threeSum(arr []int) [][]int {
	s.sortInts(arr)
	result := [][]int{}

	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left := i + 1
		right := len(arr) - 1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]

			if sum == 0 {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				left++
				right--

				for left < right && arr[left] == arr[left-1] {
					left++
				}
				for left < right && arr[right] == arr[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
