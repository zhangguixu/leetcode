package datastructure

import (
	"testing"
)

func Test_stack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	if s.Len() == 2 {
		t.Log("push success")
	} else {
		t.Error("push fail")
	}
	n := s.Pop()
	if n == 2 {
		t.Log("push success")
	} else {
		t.Error("push fail")
	}
}
