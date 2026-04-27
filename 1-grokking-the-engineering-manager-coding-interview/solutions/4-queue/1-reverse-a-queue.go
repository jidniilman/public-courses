package main

import "fmt"

func main() {
	s := Solution{}

	q1 := NewQueue()
	q1.Push(3)
	q1.Push(5)
	q1.Push(2)
	result1 := s.reverseQueue(q1)
	fmt.Print("[")
	for !result1.Empty() {
		val, _ := result1.Front()
		fmt.Print(val)
		result1.Pop()
		if !result1.Empty() {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")

	q2 := NewQueue()
	q2.Push(7)
	result2 := s.reverseQueue(q2)
	fmt.Print("[")
	for !result2.Empty() {
		val, _ := result2.Front()
		fmt.Print(val)
		result2.Pop()
		if !result2.Empty() {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")

	q3 := NewQueue()
	q3.Push(-1)
	q3.Push(0)
	q3.Push(1)
	result3 := s.reverseQueue(q3)
	fmt.Print("[")
	for !result3.Empty() {
		val, _ := result3.Front()
		fmt.Print(val)
		result3.Pop()
		if !result3.Empty() {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

type Solution struct{}

func (s *Solution) reverseQueue(queue Queue) Queue {
	stack := NewStack()

	for !queue.Empty() {
		val, _ := queue.Front()
		stack.Push(val)
		queue.Pop()
	}

	for !stack.Empty() {
		val, _ := stack.Pop()
		queue.Push(val)
	}

	return queue
}

type Queue struct {
	items []interface{}
}

func NewQueue() Queue {
	return Queue{items: []interface{}{}}
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) Push(value interface{}) {
	q.items = append(q.items, value)
}

func (q *Queue) Pop() error {
	if q.Empty() {
		return fmt.Errorf("queue is empty")
	}
	q.items = q.items[1:]
	return nil
}

type Stack struct {
	items []interface{}
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}
