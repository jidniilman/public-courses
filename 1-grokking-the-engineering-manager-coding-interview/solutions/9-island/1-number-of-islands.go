package main

// Solution struct to replace the Solution class in Java
type Solution struct{}

// countIslands method converted from Java to Go
func (s *Solution) countIslands(matrix [][]int) int {
	totalIslands := 0
	if len(matrix) == 0 {
		return totalIslands
	}

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				totalIslands++
				s.visitIslandDFS(matrix, i, j)
			}
		}
	}

	return totalIslands
}

func (s *Solution) visitIslandDFS(matrix [][]int, i, j int) {
	rows := len(matrix)
	cols := len(matrix[0])

	if i < 0 || i >= rows || j < 0 || j >= cols || matrix[i][j] == 0 {
		return
	}

	matrix[i][j] = 0

	s.visitIslandDFS(matrix, i+1, j)
	s.visitIslandDFS(matrix, i-1, j)
	s.visitIslandDFS(matrix, i, j+1)
	s.visitIslandDFS(matrix, i, j-1)
}
