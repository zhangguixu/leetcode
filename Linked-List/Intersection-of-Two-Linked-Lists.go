package main

/*
	这种方式的思路：对齐链表A，B的遍历起点

	这种实现遍历的次数和以下是一样的，但是在for循环的少比较了一次，因此执行速度是比v1版本的快
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
  var lenA, lenB int
  curA, curB := headA, headB
  for curA != nil {
    lenA++
    curA = curA.Next
  }
  for curB != nil {
    lenB++
    curB = curB.Next
  }
  curA, curB = headA, headB
  if lenA < lenB {
    curA, curB = curB, curA
    lenA, lenB = lenB, lenA
  }
  for lenA != lenB {
    curA = curA.Next
    lenA--
  }
  for curA != nil {
    if curA == curB {
      return curA
    }
    curA = curA.Next
    curB = curB.Next
  }
  return nil
}

/*
	同样是基于上述的思路，有一种更加简洁的解法：

	A和B两个链表长度可能不同，但是A+B和B+A的长度是相同的，
	所以遍历A+B和遍历B+A一定是同时结束。 
	如果A,B相交的话A和B有一段尾巴是相同的，所以两个遍历的指针一定会同时到达交点
	如果A,B不相交的话两个指针就会同时到达A+B（B+A）的尾节点
*/
func getIntersectionNode_v1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = cur.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
