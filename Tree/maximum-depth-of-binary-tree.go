package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
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
  return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}