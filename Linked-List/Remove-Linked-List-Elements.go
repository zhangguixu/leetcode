package main

/*
	这道题主要是处理一个边界的问题
	当链表所有的元素都被删除了，之后要如何返回

	技巧是使用dump node来简化边界处理的复杂度
*/
func removeElements(head *ListNode, val int) *ListNode {
    cur := &ListNode{ Next: head, }
    head = cur
    for cur.Next != nil {
        if cur.Next.Val == val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return head.Next
}