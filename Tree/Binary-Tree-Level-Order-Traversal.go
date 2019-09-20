package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func create(list []interface{}) *TreeNode {
// 	if len(list) == 0 {
// 		return nil
// 	}
// 	tnList := make([]*TreeNode, 0, 10)
// 	root := new(TreeNode)
// 	root.Val = list[0]
// 	tnList = append(tnList, root)
// 	index := 1
// 	for i := 0; i < len(tnList) && index < len(list); i++ {
// 		parent := tnList[i]
// 		n
// 		if index < len(list) {
// 			parent.Left = new(TreeNode)
// 			parent.Left.Val = list[index]
// 			tnList = append(tnList, parent.Left)
// 			index++
// 		}
// 		if index < len(list) {
// 			parent.Right = new(TreeNode)
// 			parent.Right.Val = list[index]
// 			tnList = append(tnList, parent.Right)
// 			index++
// 		}
// 	}
// 	return root
// }

func main() {
	fmt.Println()
}

func levelOrder(root *TreeNode) [][]int {
	nextQueue := make([]*TreeNode, 0, 10)
	if root != nil {
		nextQueue = append(nextQueue, root)
	}
	result := make([][]int, 0, 10)
	for len(nextQueue) > 0 {
		list := make([]int, 0, len(nextQueue))
		curQueue := nextQueue
		nextQueue = make([]*TreeNode, 0, 10)
		for _, tn := range curQueue {
			list = append(list, tn.Val)
			if tn.Left != nil {
				nextQueue = append(nextQueue, tn.Left)
			}
			if tn.Right != nil {
				nextQueue = append(nextQueue, tn.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

func levelOrder2(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0, 10)
	var curIndex, curLevelNodeNum, nextLevelNodeNum int
	if root != nil {
		queue = append(queue, root)
		nextLevelNodeNum = 1
	}
	result := make([][]int, 0, 10)
	for nextLevelNodeNum > 0 {
		list := make([]int, 0, 10)
		curLevelNodeNum = nextLevelNodeNum
		nextLevelNodeNum = 0
		for curLevelNodeNum > 0 {
			tn := queue[curIndex]
			list = append(list, tn.Val)
			if tn.Left != nil {
				queue = append(queue, tn.Left)
				nextLevelNodeNum++
			}
			if tn.Right != nil {
				queue = append(queue, tn.Right)
				nextLevelNodeNum++
			}
			curIndex++
			curLevelNodeNum--
		}
		result = append(result, list)
	}
	return result
}
