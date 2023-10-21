package queue

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		make([]int, 0, 10),
	}
}

func (q *Queue) Push(num int) {
	q.data = append(q.data, num)
}

func (q *Queue) Pop() {
	if len(q.data) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	q.data = q.data[1:]
}
