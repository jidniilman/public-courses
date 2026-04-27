package main

func main() {
}

// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

type Solution struct{}

func (s *Solution) findPath(root *TreeNode, sequence []int) bool {
	if root == nil {
		return len(sequence) == 0
	}
	return s.findPathRecursive(root, sequence, 0)
}

func (s *Solution) findPathRecursive(node *TreeNode, sequence []int, index int) bool {
	if node == nil {
		return false
	}

	if index >= len(sequence) || node.Val != sequence[index] {
		return false
	}

	if node.Left == nil && node.Right == nil && index == len(sequence)-1 {
		return true
	}

	return s.findPathRecursive(node.Left, sequence, index+1) || s.findPathRecursive(node.Right, sequence, index+1)
}
