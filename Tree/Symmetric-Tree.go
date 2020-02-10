package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历树的每一层的节点，判断是否对称
func isSymmetric(root *TreeNode) bool {
	var nextQueue []*TreeNode
	curQueue := make([]*TreeNode, 0)
	curQueue = append(curQueue, root)
	for len(curQueue) > 0 {
		nextQueue = make([]*TreeNode, 0, len(curQueue) * 2)
		for _, node := range curQueue {
			if node != nil {
				nextQueue = append(nextQueue, node.Left)
				nextQueue = append(nextQueue, node.Right)
			}
		}
		if isSymmetricList(nextQueue) {
			curQueue = nextQueue
		} else {
			return false
		}
	}
	return true
}

func isSymmetricList(list []*TreeNode) bool {
	i, j := 0, len(list) - 1
	for i < j {
		if list[i] == nil && list[j] == nil || (list[i] != nil && list[j] != nil && list[i].Val == list[j].Val) {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

// 官方的解法

/*
	递归：
	因此，该问题可以转化为：两个树在什么情况下互为镜像？

	如果同时满足下面的条件，两个树互为镜像：

	它们的两个根结点具有相同的值。
	每个树的右子树都与另一个树的左子树镜像对称。

	作者：LeetCode
	链接：https://leetcode-cn.com/problems/symmetric-tree/solution/dui-cheng-er-cha-shu-by-leetcode/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}