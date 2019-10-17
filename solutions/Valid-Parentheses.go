package main

import "fmt"

func main() {
	fmt.Println(isValid("["))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))

	// 边界
	fmt.Println(isValid("}]"))
	fmt.Println(isValid("{["))
}

type Stack struct {
	list []rune
}

func (s *Stack) Push(r rune) {
	s.list = append(s.list, r)
}

func (s *Stack) Pop() string {
	r := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return string(r)
}

func (s *Stack) Len() int {
	return len(s.list)
}

func NewStack() *Stack {
	s := new(Stack)
	s.list = make([]rune, 0)
	return s
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := NewStack()
	valid := true
	for _, r := range s {
		switch string(r) {
		case "(", "{", "[":
			stack.Push(r)
		case ")":
			if stack.Len() == 0 || stack.Pop() != "(" {
				valid = false
			}
		case "}":
			if stack.Len() == 0 || stack.Pop() != "{" {
				valid = false
			}
		case "]":
			if stack.Len() == 0 || stack.Pop() != "[" {
				valid = false
			}
		}
		if !valid {
			break
		}
	}

	if stack.Len() > 0 {
		valid = false
	}

	return valid
}
