package main

import (
	"fmt"
)

type ListNode struct {
	Val int
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
	print(removeNthFromEndByOnePass(create([]int{1,2,3,4,5,}), 2))
	print(removeNthFromEndByOnePass(create([]int{1,}), 1))
	print(removeNthFromEndByOnePass(create([]int{1,2}), 2))
}

/*
	题目简单，但是情况要想周全
	 1->2->3->4->5, and n = 2.
	 1->2->3->5

	 1->2 and n = 2
	 
	 这个解法太丑陋的了，增加一个亚节点(dump node)可以减少一次if的判断
*/
func removeNthFromEndUgly(head *ListNode, n int) *ListNode {
	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	count = count - n
	if count == 0 {
		head = head.Next
	} else {
		cur = head
		for count > 1 {
			count--
			cur = cur.Next
		}
		if cur.Next != nil {
			cur.Next = cur.Next.Next
		}
	}
	return head
}

// dump node的应用，可以减少很多不必要的判断！！
func removeNthFromEndByDumpNode(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	cur = dummy
	count -= n
	for count > 0 {
		count--
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

/*
	上面的解法需要两次循环，可以使用一次循环也可以
	技巧是使用两个指针

	fast 先前进n个节点
	之后
	slow和fast再一起前进，当fast到达链表末端的时候，这个slow和fast就相隔n个节点
	slow所在的位置就是待删除节点的上一个节点
*/
func removeNthFromEndByOnePass(head *ListNode, n int) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head

	fast := dummy
	slow := dummy

	for i := 1; i <= n + 1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	
	slow.Next = slow.Next.Next

	return dummy.Next
}

