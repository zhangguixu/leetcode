package main 

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


/*
	因为数组已经排好序，可以降低很多的复杂度，
	使用递归的思想，每次都从数组的中间挑选元素即可

	这道题 leetcode的时间判断是有问题，我提交了两次，两次的runtime相差较大

	然后找了runtime为80ms的提交了，runtime结果是125ms，我估计是测试用例有变化的原因
*/
func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums) - 1)
}

func buildBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	m := (end + start) / 2
	var left, right *TreeNode
	if start < end {
		left = buildBST(nums, start, m - 1)
		right = buildBST(nums, m + 1, end)
	}
	return &TreeNode{
		Val: nums[m],
		Left: left,
		Right: right,
	}
}