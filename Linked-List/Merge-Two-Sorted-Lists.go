package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func create(l []int) *ListNode {
	header := new(ListNode)
	header.Val = l[0]
	cur := header
	for i := 1; i < len(l); i++ {
		next := new(ListNode)
		next.Val = l[i]
		cur.Next = next
		cur = next
	}
	return header
}

func print(header *ListNode) {
	for header != nil {
		fmt.Printf(" %d ", header.Val)
		header = header.Next
	}
	fmt.Printf("\n")
}

func main() {
	print(mergeTwoLists(create([]int{1, 2, 4}), create([]int{1, 3, 4})))
	print(mergeTwoLists(create([]int{1, 2, 4}), nil))
	print(mergeTwoLists(nil, create([]int{1, 3, 4})))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var header, tail *ListNode
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		var next *ListNode
		if l1.Val <= l2.Val {
			next = l1
			l1 = l1.Next
		} else {
			next = l2
			l2 = l2.Next
		}
		if header == nil {
			header = next
		}
		if tail == nil {
			tail = next
		} else {
			tail.Next = next
			tail = tail.Next
		}
	}
	for l1 != nil {
		tail.Next = l1
		tail = tail.Next
		l1 = l1.Next
	}
	for l2 != nil {
		tail.Next = l2
		tail = tail.Next
		l2 = l2.Next
	}
	return header
}
