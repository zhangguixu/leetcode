package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}
	dumpNode := &ListNode{Next: head}
	pre := dumpNode
	var rHead, rTail, next *ListNode
	for head != nil && m > 1 {
		m--
		n--
		pre = head
		head = head.Next
	}
	rTail = head
	for head != nil && n > 0 {
		n--
		next = head.Next
		head.Next = rHead
		rHead = head
		head = next
	}
	pre.Next = rHead
	rTail.Next = head
	return dumpNode.Next
}