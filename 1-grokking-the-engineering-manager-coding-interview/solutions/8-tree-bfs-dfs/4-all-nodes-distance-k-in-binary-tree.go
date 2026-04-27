package main

func main() {
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// // Custom Queue implementation
// type Queue struct {
// 	data *list.List
// }

// func NewQueue() Queue {
// 	return Queue{data: list.New()}
// }

// func (q *Queue) Empty() bool {
// 	return q.data.Len() == 0
// }

// func (q *Queue) Front() (interface{}, error) {
// 	if q.Empty() {
// 		return nil, errors.New("queue is empty")
// 	}
// 	return q.data.Front().Value, nil
// }

// func (q *Queue) Push(value interface{}) {
// 	q.data.PushBack(value)
// }

// func (q *Queue) Pop() error {
// 	if q.Empty() {
// 		return errors.New("queue is empty")
// 	}
// 	q.data.Remove(q.data.Front())
// 	return nil
// }

type Solution struct{}

func (s *Solution) distanceK(root *TreeNode, target *TreeNode, K int) []int {
	result := make([]int, 0)
	parentMap := make(map[*TreeNode]*TreeNode)
	s.buildParentMap(root, nil, parentMap)

	visited := make(map[*TreeNode]bool)
	s.findNodesAtDistanceK(target, K, parentMap, visited, &result)
	return result
}

func (s *Solution) buildParentMap(node *TreeNode, parent *TreeNode, parentMap map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}
	parentMap[node] = parent
	s.buildParentMap(node.Left, node, parentMap)
	s.buildParentMap(node.Right, node, parentMap)
}

func (s *Solution) findNodesAtDistanceK(node *TreeNode, k int, parentMap map[*TreeNode]*TreeNode, visited map[*TreeNode]bool, result *[]int) {
	if node == nil || visited[node] {
		return
	}

	visited[node] = true

	if k == 0 {
		*result = append(*result, node.Val)
		return
	}

	s.findNodesAtDistanceK(node.Left, k-1, parentMap, visited, result)
	s.findNodesAtDistanceK(node.Right, k-1, parentMap, visited, result)
	s.findNodesAtDistanceK(parentMap[node], k-1, parentMap, visited, result)
}
