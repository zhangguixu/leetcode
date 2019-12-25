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
	print(partition(create([]int{1,4,3,2,5,2}), 3))
}

// 分区操作，是快排的基础，在此基础可以考察链表的快排实现
func partition(head *ListNode, x int) *ListNode {
	lHead, rHead := new(ListNode), new(ListNode)
	lTail, rTail := lHead, rHead
	for head != nil {
		if head.Val < x {
			lTail.Next = head
			lTail = head
		} else {
			rTail.Next = head
			rTail = head
		}
		head = head.Next
	}
	// 注意，如果不置为nil的话，可以会形成环，导致最终结果over time limit
	rTail.Next = nil
	lTail.Next = rHead.Next
	return lHead.Next
}