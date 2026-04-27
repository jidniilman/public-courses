package main

import "fmt"

func main() {
	s := Solution{}
	rooms1 := [][]int{
		{2147483647, -1, 0},
		{2147483647, 2147483647, 2147483647},
		{2147483647, -1, 2147483647},
	}
	s.wallsAndGates(rooms1)
	fmt.Println(rooms1)

	rooms2 := [][]int{
		{0, 2147483647, 2147483647, 0},
		{2147483647, -1, 2147483647, 2147483647},
		{2147483647, -1, -1, 2147483647},
		{0, 2147483647, 2147483647, 0},
	}
	s.wallsAndGates(rooms2)
	fmt.Println(rooms2)
}

type Solution struct{}

type Cell struct {
	row, col int
}

func (s *Solution) wallsAndGates(rooms [][]int) [][]int {
	if len(rooms) == 0 {
		return rooms
	}

	rows := len(rooms)
	cols := len(rooms[0])
	queue := []Cell{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, Cell{i, j})
			}
		}
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow := cell.row + dir[0]
			newCol := cell.col + dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && rooms[newRow][newCol] == 2147483647 {
				rooms[newRow][newCol] = rooms[cell.row][cell.col] + 1
				queue = append(queue, Cell{newRow, newCol})
			}
		}
	}

	return rooms
}
