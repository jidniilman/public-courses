package main

import "fmt"

func main() {
	s := Solution{}
	grid := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println(s.maxAreaOfIsland(grid))
}

type Solution struct{}

func (s *Solution) maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	maxArea := 0
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				area := s.dfs(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func (s *Solution) dfs(grid [][]int, i, j int) int {
	rows := len(grid)
	cols := len(grid[0])

	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	area := 1
	area += s.dfs(grid, i+1, j)
	area += s.dfs(grid, i-1, j)
	area += s.dfs(grid, i, j+1)
	area += s.dfs(grid, i, j-1)

	return area
}
