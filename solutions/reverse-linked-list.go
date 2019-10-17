package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Print() {
	list := make([]int, 0)
	cur := head
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}
	fmt.Println(list)
}

func (head *ListNode) Insert(val int) {
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = new(ListNode)
	cur.Next.Val = val
}

func main() {
	head := new(ListNode)
	head.Val = 1

	// for i := 2; i <= 8; i++ {
	// 	head.Insert(i)
	// }

	// head.Print()
	// reverseList(head)
	reverseList(head).Print()
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode
	var next *ListNode

	cur = head

	for cur != nil {
		next = cur.Next
		fmt.Printf("pre：%v,cur: %v, next: %v \n", pre, cur, next)
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
