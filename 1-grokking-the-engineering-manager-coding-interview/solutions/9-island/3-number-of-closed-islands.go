package main

import "fmt"

func main() {
	s := Solution{}
	grid1 := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	fmt.Println(s.closedIsland(grid1))

	grid2 := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
	}
	fmt.Println(s.closedIsland(grid2))
}

type Solution struct{}

func (s *Solution) closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		if grid[i][0] == 1 {
			s.dfs(grid, i, 0)
		}
		if grid[i][cols-1] == 1 {
			s.dfs(grid, i, cols-1)
		}
	}

	for j := 0; j < cols; j++ {
		if grid[0][j] == 1 {
			s.dfs(grid, 0, j)
		}
		if grid[rows-1][j] == 1 {
			s.dfs(grid, rows-1, j)
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				count++
				s.dfs(grid, i, j)
			}
		}
	}

	return count
}

func (s *Solution) dfs(grid [][]int, i, j int) {
	rows := len(grid)
	cols := len(grid[0])

	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0

	s.dfs(grid, i+1, j)
	s.dfs(grid, i-1, j)
	s.dfs(grid, i, j+1)
	s.dfs(grid, i, j-1)
}
