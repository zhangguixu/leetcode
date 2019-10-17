package datastructure

/*
	二叉查找树，wiki定义：

	1. 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
	2. 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
	3. 任意节点的左、右子树也分别为二叉查找树；
	4. 没有键值相等的节点。

	二叉查找树相比于其他数据结构的优势在于查找、插入的时间复杂度较低。为O(logn)。
	二叉查找树是基础性数据结构，用于构建更为抽象的数据结构，如集合、多重集、关联数组等。

	详细 https://zh.wikipedia.org/wiki/%E4%BA%8C%E5%85%83%E6%90%9C%E5%B0%8B%E6%A8%B9
*/

// type TreeNode struct {
// 	Val int
// 	Left *TreeNode
// 	Right *TreeNode
// }

// func insert(t *TreeNode, val int) {
// 	if t == nil {
// 		t = new(TreeNode)
// 		t.Val = val
// 		return
// 	}
// 	if val <= t.Val {
// 		insert(t.Left, val)
// 	} else {
// 		insert(t.Right, val)
// 	}
// }

// func (t *TreeNode) Insert(val int) {
// 	insert(t, val)
// }
