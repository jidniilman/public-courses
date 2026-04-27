package main

import "fmt"

func main() {
	s := Solution{}
	arr1 := []int{1, 0, 2, 1, 0}
	s.sortColors(arr1)
	fmt.Println(arr1)

	arr2 := []int{2, 2, 0, 1, 2, 0}
	s.sortColors(arr2)
	fmt.Println(arr2)
}

type Solution struct{}

func (s *Solution) sortColors(arr []int) []int {
	low := 0
	mid := 0
	high := len(arr) - 1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}

	return arr
}
