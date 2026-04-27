package main

import "fmt"

func main() {
	s := Solution{}
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	fmt.Println(s.isEvenOddTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct{}

func (s *Solution) isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		prevVal := -1

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				if node.Val%2 == 0 {
					return false
				}
				if prevVal != -1 && node.Val <= prevVal {
					return false
				}
			} else {
				if node.Val%2 == 1 {
					return false
				}
				if prevVal != -1 && node.Val >= prevVal {
					return false
				}
			}

			prevVal = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}

	return true
}
