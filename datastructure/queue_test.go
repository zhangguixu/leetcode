package datastructure

import (
	"testing"
)

func Test_queue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	if q.Len() == 2 {
		t.Log("enqueue success")
	} else {
		t.Error("enqueue fail")
	}
	n := q.Dequeue()
	if n == 1 {
		t.Log("dequeue success")
	} else {
		t.Error("dequeue fail")
	}
}
