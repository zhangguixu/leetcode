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
}

func isPalindrome(head *ListNode) bool {
	list := make([]int, 0, 10)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	m := len(list) / 2
	start := 0
	end := len(list) - 1
	flag := true
	for start < m {
		if list[start] != list[end] {
			flag = false
			break
		}
		start++
		end--
	}
	return flag
}
