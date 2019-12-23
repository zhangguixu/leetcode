package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func create(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}
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
	print(rotateRight(create([]int{1,2}), 1)) // 2, 1
	print(rotateRight(create([]int{1,2,3,4,5}), 2)) // 4,5,1,2,3
	print(rotateRight(create([]int{1}), 0)) // 1
	print(rotateRight(create([]int{1}), 0)) // 1
	print(rotateRight(create([]int{}), 0)) // 
	print(rotateRight(create([]int{0,1,2}), 4)) // 1
}

// 这道题主要还是考察一个对题目的理解，特别注意k = total- k % total - 1
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var tail *ListNode
	total, cur := 0, head
	for cur != nil {
		total++
		tail = cur
		cur = cur.Next
	}
	k = total - k % total - 1
	cur = head
	for k > 0 {
		cur = cur.Next
		k--
	}
	tail.Next = head
	head = cur.Next
	cur.Next = nil
	return head
}