package main

func main() {

}

// Queue structure using list
// type Queue struct {
// 	data *list.List
// }

// // NewQueue initializes a new queue
// func NewQueue() Queue {
// 	return Queue{data: list.New()}
// }

// // Empty checks if the queue is empty
// func (q *Queue) Empty() bool {
// 	return q.data.Len() == 0
// }

// // Front gets the front element without removing it
// func (q *Queue) Front() (interface{}, error) {
// 	if q.Empty() {
// 		return nil, errors.New("queue is empty")
// 	}
// 	return q.data.Front().Value, nil
// }

// // Push adds an element to the back of the queue
// func (q *Queue) Push(value interface{}) {
// 	q.data.PushBack(value)
// }

// // Pop removes the front element from the queue
// func (q *Queue) Pop() error {
// 	if q.Empty() {
// 		return errors.New("queue is empty")
// 	}
// 	q.data.Remove(q.data.Front())
// 	return nil
// }

// Solution struct
type Solution struct {
	queue1 []int
	queue2 []int
}

// Constructor initializes the Solution struct
func constructor() *Solution {
	return &Solution{
		queue1: []int{},
		queue2: []int{},
	}
}

// Push element x onto the stack
func (this *Solution) push(x int) {
	this.queue1 = append(this.queue1, x)
}

// Pop removes and returns the top element from the stack
func (this *Solution) pop() int {
	for len(this.queue1) > 1 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}

	result := this.queue1[0]
	this.queue1 = this.queue2
	this.queue2 = []int{}

	return result
}

// Top returns the top element without removing it
func (this *Solution) top() int {
	for len(this.queue1) > 1 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}

	result := this.queue1[0]
	this.queue2 = append(this.queue2, result)
	this.queue1 = this.queue2
	this.queue2 = []int{}

	return result
}

// Empty checks if the stack is empty
func (this *Solution) empty() bool {
	return len(this.queue1) == 0
}
