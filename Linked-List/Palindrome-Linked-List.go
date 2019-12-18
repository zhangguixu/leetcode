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

func main() {
	fmt.Println(isPalindrome(create([]int{1, 2})))
	fmt.Println(isPalindrome(create([]int{1, 2, 3})))
	fmt.Println(isPalindrome(create([]int{1, 2, 3, 3, 2, 2})))
	fmt.Println(isPalindrome(create([]int{1, 0, 1})))
	fmt.Println(isPalindrome(create([]int{1})))
	fmt.Println(isPalindrome(create([]int{})))

	fmt.Println("====================================")

	fmt.Println(isPalindrome_v1(create([]int{1, 2})))
	fmt.Println(isPalindrome_v1(create([]int{1, 2, 3})))
	fmt.Println(isPalindrome_v1(create([]int{1, 2, 3, 3, 2, 2})))
	fmt.Println(isPalindrome_v1(create([]int{1, 0, 1})))
	fmt.Println(isPalindrome_v1(create([]int{1})))
	fmt.Println(isPalindrome_v1(create([]int{})))
}

// 借助O(n)的空间
func isPalindrome(head *ListNode) bool {
	list := make([]int, 0, 10)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	l, m, r := 0, len(list) / 2, len(list) - 1
	for l < m {
		if list[l] != list[r] {
			return false
		}
		l++
		r--
	}
	return true
}

/*
	时间复杂度：O(n)
	空间复杂度：O(1)

	涉及到两个技巧
	1）使用快慢双指针找到链表的中间节点
	2）链表的反转

	这种实现会破坏原来的链表结构，因此fast的移动必须要先与slow
*/
func isPalindrome_v1(head *ListNode) bool {
	var pre, next *ListNode
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			slow = slow.Next
			break
		}
		fast = fast.Next.Next
		next = slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	for pre != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}

/*
	快慢指针找到中间节点

	加上栈（存储空间比第一种时间少一半）
*/
