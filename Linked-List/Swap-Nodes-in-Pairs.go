package main

/*
	链表的题目可以通过画图来简化思考的复杂度
	链表的交换还是比数组的复杂很多
*/
func swapPairs(head *ListNode) *ListNode {
  dumpNode := &ListNode{Next: head}
  pre, cur := dumpNode, head
  for cur != nil && cur.Next != nil {
    pre.Next = cur.Next
    cur.Next = cur.Next.Next
    pre.Next.Next = cur
    pre = cur
    cur = cur.Next
  }
  return dumpNode.Next
}