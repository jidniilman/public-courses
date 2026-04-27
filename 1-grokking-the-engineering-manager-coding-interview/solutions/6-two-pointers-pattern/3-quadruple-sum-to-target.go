package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.fourSum([]int{4, 1, 2, -1, 1, -3}, 1))
	fmt.Println(s.fourSum([]int{2, 0, -1, 1, -2, 2}, 2))
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

func (s *Solution) fourSum(arr []int, target int) [][]int {
	s.sortInts(arr)
	result := [][]int{}
	n := len(arr)

	for i := 0; i < n-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			left := j + 1
			right := n - 1

			for left < right {
				sum := arr[i] + arr[j] + arr[left] + arr[right]

				if sum == target {
					result = append(result, []int{arr[i], arr[j], arr[left], arr[right]})
					left++
					right--

					for left < right && arr[left] == arr[left-1] {
						left++
					}
					for left < right && arr[right] == arr[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}
