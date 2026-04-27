package main

import "fmt"

func main() {
	fmt.Println(transformArray([]int{9, 9, 9, 9, 9}))
}

func transformArray(arr []int) []int {
	for {
		isChanged := false
		for i := 0; i < len(arr); i++ {
			if i != 0 && i != len(arr)-1 {
				if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
					arr[i]--
					isChanged = true
				} else if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
					arr[i]++
					isChanged = true
				} else if arr[i] == arr[i-1] && arr[i] == arr[i+1] {
					isChanged = false
				}
			}
		}
		if !isChanged {
			return arr
		}
	}
}
