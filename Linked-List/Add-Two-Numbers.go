package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func create(l []int) *ListNode {
	header := new(ListNode)
	if len(l) == 0 {
		return nil
	}
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
	print(addTwoNumbers(create([]int{1,2}), create([]int{})))
	print(addTwoNumbers(create([]int{}), create([]int{1,2})))
	print(addTwoNumbers(create([]int{8,9,9}), create([]int{2,2,2})))
	print(addTwoNumbers(create([]int{2,4,3}), create([]int{5,6,4})))
	print(addTwoNumbers(create([]int{8,9,9}), create([]int{2,2,2,1})))
	print(addTwoNumbers(create([]int{2,2,2,1}), create([]int{8,9,9})))
	print(addTwoNumbers(create([]int{8,9,9}), create([]int{2})))
	print(addTwoNumbers(create([]int{2}), create([]int{8,9,9})))
}

/*
	本题的难点在于：边界条件比较多，需要考虑几种情况，不然很容易出现bad case
	
	还有就是利用原来链表的空间（当然其实是对原有链表的破坏），减少空间的开销
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre *ListNode
	var carry int
	cur := l1
	for cur != nil && l2 != nil {
		cur.Val += l2.Val + carry
		carry = 0
		if cur.Val >= 10 {
			cur.Val -= 10
			carry = 1
		}
		pre = cur
		cur = cur.Next
		l2 = l2.Next
	}
	if l2 != nil {
		if pre == nil {
			l1 = l2
		} else {
			pre.Next = l2
			cur = l2
		}
	}
	for cur != nil {
		cur.Val += carry
		if cur.Val < 10 {
			carry = 0
			break
		}
		cur.Val -= 10
		pre = cur
		cur = cur.Next
	}
	if carry == 1 && cur == nil {
		pre.Next = &ListNode{Val: 1}
	}
	return l1
}