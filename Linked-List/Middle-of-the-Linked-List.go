package main

// 快慢指针法
func middleNode(head *ListNode) *ListNode {
  slow, fast := head, head
  for fast != nil && fast.Next != nil  {
    slow = slow.Next
    fast = fast.Next.Next
  }
  return slow
}