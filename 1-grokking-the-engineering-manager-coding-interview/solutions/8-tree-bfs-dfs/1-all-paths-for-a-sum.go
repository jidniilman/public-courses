package main

func main() {

}

// TreeNode struct
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// Solution struct
type Solution struct{}

// FindPaths to find paths that sum to a target value
func (s *Solution) findPaths(root *TreeNode, sum int) [][]int {
	allPaths := [][]int{}
	if root == nil {
		return allPaths
	}
	s.findPathsRecursive(root, sum, []int{}, &allPaths)
	return allPaths
}

func (s *Solution) findPathsRecursive(node *TreeNode, sum int, currentPath []int, allPaths *[][]int) {
	if node == nil {
		return
	}

	currentPath = append(currentPath, node.Val)

	if node.Left == nil && node.Right == nil && node.Val == sum {
		path := make([]int, len(currentPath))
		copy(path, currentPath)
		*allPaths = append(*allPaths, path)
	} else {
		s.findPathsRecursive(node.Left, sum-node.Val, currentPath, allPaths)
		s.findPathsRecursive(node.Right, sum-node.Val, currentPath, allPaths)
	}
}
