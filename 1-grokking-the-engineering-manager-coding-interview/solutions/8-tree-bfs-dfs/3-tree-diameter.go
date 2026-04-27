package main

func main() {
}

// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

type Solution struct {
	treeDiameter int
}

func (s *Solution) findDiameter(root *TreeNode) int {
	s.treeDiameter = 0
	s.calculateHeight(root)
	return s.treeDiameter
}

func (s *Solution) calculateHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := s.calculateHeight(node.Left)
	rightHeight := s.calculateHeight(node.Right)

	currentDiameter := leftHeight + rightHeight + 1
	if currentDiameter > s.treeDiameter {
		s.treeDiameter = currentDiameter
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
