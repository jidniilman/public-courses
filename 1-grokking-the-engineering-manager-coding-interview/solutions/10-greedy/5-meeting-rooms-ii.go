package main

import (
	"fmt"
	"sort"
)

func main() {
	s := Solution{}
	fmt.Println(s.minMeetingRooms([][]int{{10, 15}, {20, 25}, {30, 35}}))
	fmt.Println(s.minMeetingRooms([][]int{{10, 20}, {15, 25}, {24, 30}, {5, 14}, {22, 28}, {1, 4}, {27, 35}}))
	fmt.Println(s.minMeetingRooms([][]int{{10, 20}, {20, 30}}))
}

type Solution struct{}

func (s *Solution) minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	rooms := 0
	endPointer := 0

	for i := 0; i < len(starts); i++ {
		if starts[i] < ends[endPointer] {
			rooms++
		} else {
			endPointer++
		}
	}

	return rooms
}
