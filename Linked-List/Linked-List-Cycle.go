package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(l []int, pos int) *ListNode {
	var header, cur, target *ListNode

	for i := 0; i < len(l); i++ {
		next := new(ListNode)
		next.Val = l[i]
		if header == nil {
			header = next
			cur = next
		} else {
			cur.Next = next
			cur = next
		}
		if pos == i {
			target = next
		}
	}

	if target != nil {
		cur.Next = target
	}

	return header
}

func print(header *ListNode, total int) {
	for header != nil && total > 0 {
		fmt.Printf(" %d ", header.Val)
		header = header.Next
		total--
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println(hasCycle(createList([]int{3, 2, 0, -4}, 1)))
}

/*
	判断链表是否有环

	这里的技巧就是使用快慢指针

	定义两个指针，同时从链表的头节点出发，

	一个指针一次走一步，
	另一个指针一次走两步。

	如果走得快的指针追上了走得慢的指针，那么链表就是环形链表；

	如果走得快的指针走到了链表的末尾（next指向 NULL）都没有追上第一个指针，那么链表就不是环形链表。

*/
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	cycleF := false
	for fast != nil && fast.Next != nil {
		if fast.Next == slow || fast.Next.Next == slow {
			cycleF = true
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return cycleF
}
