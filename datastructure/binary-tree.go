package datastructure

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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preVisit(t *TreeNode) []int {
	order := make([]int, 0)
	if t != nil {
		order = append(order, t.Val)
		lOrder := preVisit(t.Left)
		rOrder := preVisit(t.Right)
		order = append(order, lOrder...)
		order = append(order, rOrder...)
	}
	return order
}

func inVisit(t *TreeNode) []int {
	order := make([]int, 0)

	if t != nil {
		lOrder := inVisit(t.Left)
		order = append(order, lOrder...)
		order = append(order, t.Val)
		rOrder := inVisit(t.Right)
		order = append(order, rOrder...)
	}

	return order
}

func postVisit(t *TreeNode) []int {
	order := make([]int, 0)
	if t != nil {
		lOrder := postVisit(t.Left)
		order = append(order, lOrder...)
		rOrder := postVisit(t.Right)
		order = append(order, rOrder...)
		order = append(order, t.Val)
	}
	return order
}

func (head *TreeNode) PreOrderVisit() []int {
	return preVisit(head)
}

func (head *TreeNode) InOrderVisit() []int {
	return inVisit(head)
}

func (head *TreeNode) PostOrderVisit() []int {
	return postVisit(head)
}

// func bfs(root *TreeNode) []int {
// 	q := NewQueue()
// 	nodes := make([]int, 0, 10)
// 	q.Enqueue(root)
// 	for q.Len() > 0 {
// 		cur := q.Dequeue()
// 		if cur != nil {
// 			nodes = append(nodes, cur.Val)
// 		}
// 		if cur.Left != nil {
// 			q.Enqueue(cur.Left)
// 		}
// 		if cur.Right != nil {
// 			q.Enqueue(cur.Right)
// 		}
// 	}
// 	return nodes
// }
