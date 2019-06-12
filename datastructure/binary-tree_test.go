package datastructure

import (
	"testing"
)

/*
	-3
		-1
			-null
			-2
		-5
			-4
			-6
	树形结构： https://images0.cnblogs.com/i/497634/201403/270932554522177.jpg
*/
func Test_preOrderVisit(t *testing.T) {
	head := createTree()
	t.Log(head.PreOrderVisit())
}

func Test_inOrderVisit(t *testing.T) {
	head := createTree()
	t.Log(head.InOrderVisit())
}

func Test_postOrderVisit(t *testing.T) {
	head := createTree()
	t.Log(head.PostOrderVisit())
}


func createTree() *TreeNode {
	head := new(TreeNode)
	head.Val = 3

	t1 := new(TreeNode)
	t1.Val = 1

	t2 := new(TreeNode)
	t2.Val = 5

	t3 := new(TreeNode)
	t3.Val = 2

	t4 := new(TreeNode)
	t4.Val = 4

	t5 := new(TreeNode)
	t5.Val = 6

	head.Left = t1
	head.Right = t2

	t1.Right = t3
	
	t2.Left = t4
	t2.Right = t5

	return head
}