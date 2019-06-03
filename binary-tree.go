package main

import "fmt"

/*
	深度优先遍历(DFS)，细分有三种

		1）pre-order前序遍历 先根后左再右
		2）in-order中序遍历，先左后中再右
		3）post-order后序遍历，先左后右再中

		深度优先的访问是栈的思想

	广度优先遍历(BFS)

		从根节点出发，沿着树的宽度访问

		广度优先的访问是队列的思想
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {

}

func bfs(root *TreeNode) []int {
}

