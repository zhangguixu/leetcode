package main 

import (
	"fmt"
	"errors"
)

type Queue struct {
	list []int
}

func NewQueue() *Queue {
	list := make([]int, 0)
	return &Queue{list}
}

func (q *Queue) Enqueue(val int) {
	q.list = append(q.list, val)
}

func (q *Queue) Dequeue() int {
	if q.Len() == 0 {
		panic(errors.New("queue is empty"))
	}
	val := q.list[0]
	q.list = q.list[1:]
	return val
}

func (q *Queue) Len() int {
	return len(q.list)
}

func main() {
	q := NewQueue()
	for i := 0; i < 9; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 9; i++ {
		fmt.Printf(" %d ", q.Dequeue())
	}
}