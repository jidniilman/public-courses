import (
	"container/list"
)

type Solution struct {
	queue *list.List
}

func Constructor(v1 []int, v2 []int) Solution {
	queue := list.New()

	maxLen := len(v1)
	if len(v2) > maxLen {
		maxLen = len(v2)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(v1) {
			queue.PushBack(v1[i])
		}
		if i < len(v2) {
			queue.PushBack(v2[i])
		}
	}

	return Solution{queue}
}

func (s *Solution) Next() int {
	if s.queue.Len() == 0 {
		return 0
	}

	front := s.queue.Front()
	s.queue.Remove(front)
	return front.Value.(int)
}

func (s *Solution) HasNext() bool {
	return s.queue.Len() > 0
}

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