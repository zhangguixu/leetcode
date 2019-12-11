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

// 这个版本的代码相当不简洁
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
  if l1 == nil {
    return l2
  }
  if l2 == nil {
    return l1
  }
  var head, tail *ListNode
  if l1.Val <= l2.Val {
    head = l1
    l1 = l1.Next
  } else {
    head = l2
    l2 = l2.Next
  }
  tail = head
  for l1 != nil && l2 != nil {
    if l1.Val <= l2.Val {
      tail.Next = l1
      l1 = l1.Next
    } else {
      tail.Next = l2
      l2 = l2.Next
    }
    tail = tail.Next
  }
  if l1 != nil {
    tail.Next = l1
  }
  if l2 != nil {
    tail.Next = l2
  }
  return head
}

// 使用(dump node)哑节点简化代码实现
func mergeTwoLists_v1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
  for l1 != nil && l2 != nil {
    if l1.Val <= l2.Val {
      tail.Next = l1
      l1 = l1.Next
    } else {
      tail.Next = l2
      l2 = l2.Next
    }
    tail = tail.Next
  }
  if l1 != nil {
    tail.Next = l1
  }
  if l2 != nil {
    tail.Next = l2
  }
  return head.Next
}
