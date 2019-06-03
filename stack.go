package main

import (
	"fmt"
	"errors"
)

type Stack struct {
	list []int
}

func NewStack() *Stack {
	list := make([]int, 0)
	return &Stack{list}
}

func (s *Stack) Len() int {
	return len(s.list)
}

func (s *Stack) Push(val int) {
	s.list = append(s.list, val)
}

func (s *Stack) Pop() int {
	if s.Len() == 0 {
		panic(errors.New("stack is empty"))
	}
	l := s.Len() - 1
	val := s.list[l]
	s.list = s.list[:l]
	return val
}

func main() {
	s := NewStack()
	for i := 0; i < 9; i++ {
		s.Push(i)
	}
	for i := 0; i < 9; i++ {
		fmt.Println(s.Pop())
	}
}