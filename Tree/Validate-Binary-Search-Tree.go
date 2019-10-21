package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// https://golangtc.com/t/5695ec6fb09ecc083100002b 
// 在go中如何模拟Integer，借鉴 sql中的实现
type NullInt struct {
	Val int
	Valid bool
}

/*
	BST(binary search tree)
	左边所有子节点都小于当前节点的值
	右边所有的子节点都大于当前节点的值

  使用递归的思想，在局部情况，待判断的节点是o3，满足BST的要求，需要满足

	  o1
	o2
	  o3     o1 > o3 > o2 

	 o1
		o2
	o3       o2 > o3 > o1

*/
func isValidBST(root *TreeNode) bool {
	// 如果不想使用的NullInt的方式，可以使用的math.MaxInt64和 math.MinInt64
	return isValid(root, NullInt{}, NullInt{})
}

func isValid(t *TreeNode, lower NullInt, upper NullInt) bool {
	if t == nil {
		return true
	}
	if lower.Valid && t.Val <= lower.Val {
		return false
	}
	if upper.Valid && t.Val >= upper.Val {
		return false
	}
	return isValid(t.Right, NullInt{t.Val,true,}, upper) && isValid(t.Left, lower, NullInt{t.Val,true,})
}
