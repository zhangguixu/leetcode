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
	需要注意的边界是删除表头的情况，按照套路就是运用 dump node
	
	dump node： 涉及表头边界逻辑一般都可以运用，简化代码逻辑分支
	快慢指针（其实就是双指针，跟数组的双指针有点像）
		中间节点
		从链表尾部的。。
		是否有环

	比较常规的套路题
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
  dumpNode := &ListNode{Next: head}
  slow, fast := dumpNode, dumpNode
  for n > 0 {
    fast = fast.Next
    n--
  }
  for fast.Next != nil {
    slow = slow.Next
    fast = fast.Next
  }
  slow.Next = slow.Next.Next
  return dumpNode.Next
}

