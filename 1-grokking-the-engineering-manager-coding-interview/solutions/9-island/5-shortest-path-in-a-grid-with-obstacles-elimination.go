package main

import "fmt"

func main() {
	s := Solution{}
	grid1 := [][]int{
		{0, 1, 0, 0},
		{1, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 1, 0},
	}
	fmt.Println(s.shortestPath(grid1, 1))

	grid2 := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(s.shortestPath(grid2, 2))

	grid3 := [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
		{0, 1, 0},
	}
	fmt.Println(s.shortestPath(grid3, 3))
}

type Solution struct{}

type State struct {
	row, col, obstacles, steps int
}

func (s *Solution) shortestPath(grid [][]int, k int) int {
	rows := len(grid)
	cols := len(grid[0])

	if rows == 1 && cols == 1 {
		return 0
	}

	queue := []State{{0, 0, 0, 0}}
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d,%d", 0, 0, 0)] = true

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow := state.row + dir[0]
			newCol := state.col + dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				newObstacles := state.obstacles + grid[newRow][newCol]

				if newObstacles <= k {
					if newRow == rows-1 && newCol == cols-1 {
						return state.steps + 1
					}

					key := fmt.Sprintf("%d,%d,%d", newRow, newCol, newObstacles)
					if !visited[key] {
						visited[key] = true
						queue = append(queue, State{newRow, newCol, newObstacles, state.steps + 1})
					}
				}
			}
		}
	}

	return -1
}
