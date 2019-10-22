package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// 这里可以模拟一个队列来做的话，代码的可读性会高一点

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftList := make([]*TreeNode, 0, 10)
	rightList := make([]*TreeNode, 0, 10)
	leftList = append(leftList, root.Left)
	rightList = append(rightList, root.Right)
	flag := true
	for len(leftList) > 0 {
		curLen := len(leftList)
		for i := 0; i < curLen; i++ {
			if leftList[i] != nil {
				leftList = append(leftList, leftList[i].Left)
				leftList = append(leftList, leftList[i].Right)
			}
			if rightList[i] != nil {
				rightList = append(rightList, rightList[i].Right)
				rightList = append(rightList, rightList[i].Left)
			}
			if leftList[i] == nil && rightList[i] == nil {
				continue
			}
			if leftList[i] != nil && rightList[i] != nil && leftList[i].Val == rightList[i].Val {
				continue
			}
			flag = false
			break			
		}
		if !flag {
			break
		}
		leftList = leftList[curLen:]
		rightList = rightList[curLen:]
		if len(leftList) != len(rightList) {
			flag = false
			break
		}
	}
	return flag
}