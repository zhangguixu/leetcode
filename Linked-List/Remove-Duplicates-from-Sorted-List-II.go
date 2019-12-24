package main

func deleteDuplicates(head *ListNode) *ListNode {
  dumpNode := &ListNode{Next: head}
  pre := dumpNode
  needDel := false
  for head != nil {
    for head.Next != nil && head.Val == head.Next.Val {
      needDel = true
      head.Next = head.Next.Next
    }
    if needDel {
      pre.Next = head.Next
      needDel = false
    } else {
      pre = head
    }
    head = head.Next
  }
  return dumpNode.Next
}