package main

// 使用两个数组交替存储当前层级和下一层级
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

// 使用一个数组存储所有的节点
func levelOrder(root *TreeNode) [][]int {
  result := make([][]int, 0)
  nodes := make([]*TreeNode, 0)
  if root != nil {
    nodes = append(nodes, root)
  }
  i, boundary := 0, len(nodes)
  var n *TreeNode
  for i < boundary {
    list := make([]int, 0)
    for i < boundary {
      n = nodes[i]
      list = append(list, n.Val)
      if n.Left != nil {
        nodes = append(nodes, n.Left)
      }
      if n.Right != nil {
        nodes = append(nodes, n.Right)
      }
      i++
    }
    result = append(result, list)
    boundary = len(nodes)
  }
  return result
}
