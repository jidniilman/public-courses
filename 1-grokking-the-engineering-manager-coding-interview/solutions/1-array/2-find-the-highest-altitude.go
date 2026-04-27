package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{2, 2, -3, -1, 2, 1, -5}))
}

func largestAltitude(gain []int) int {
	maxAltitude := 0
	result := make([]int, len(gain))
	for i := 0; i < len(gain); i++ {
		result[i] = gain[i]
		if i > 0 {
			result[i] += result[i-1]
		}
		if result[i] > maxAltitude {
			maxAltitude = result[i]
		}
	}
	return maxAltitude
}
