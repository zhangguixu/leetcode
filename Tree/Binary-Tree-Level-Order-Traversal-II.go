package main

/*
	这道题跟102的解法是一样的：广度优先 + 数组的反转
*/

/*
	这里采用递归，需要记录当前的层级，最后进行数组的反转
*/
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	result = dfs(root, 0, result)
	reverse(result)
	return result
}

func dfs(t *TreeNode, level int, result [][]int) [][]int {
	if t == nil {
		return result
	}
	for len(result) <= level {
		result = append(result, []int{})
	}
	result[level] = append(result[level], t.Val)
	result = dfs(t.Left, level + 1, result)
	result = dfs(t.Right, level + 1, result)
}

func reverse(result [][]int) {
	i, j := 0, len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
}
