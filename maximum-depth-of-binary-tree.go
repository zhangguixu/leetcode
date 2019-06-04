package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	// [3,9,20,null,null,15,7]
	r := new(TreeNode)
	r.Val = 3
	r.Left = new(TreeNode)
	r.Left.Val = 9
	r.Right = new(TreeNode)
	r.Right.Val = 20
	
	r1 := r.Right
	r1.Left = &TreeNode{15, nil, nil}
	r1.Right = &TreeNode{7, nil, nil}

	fmt.Println(maxDepth(r))
}

/*
	递归的思想
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
    return depth(root)
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func depth(t *TreeNode) int {
	lDepth := 0
	rDepth := 0
	if t.Left != nil {
		lDepth = depth(t.Left)
	}
	if t.Right != nil {
		rDepth = depth(t.Right)
	}
	return 1 + max(lDepth, rDepth)
}