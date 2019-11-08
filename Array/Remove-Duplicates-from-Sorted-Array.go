package main

import "fmt"

func print(nums []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf(" %d ", nums[i])
	}
	fmt.Printf("\n")
}

func main() {
	t1 := []int{1, 1, 2}
	print(t1, removeDuplicates(t1))

	t2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	print(t2, removeDuplicates(t2))

	t3 := []int{1, 1, 1}
	print(t3, removeDuplicates(t3))

	t4 := []int{1}
	print(t4, removeDuplicates(t4))

	t5 := []int{}
	print(t5, removeDuplicates(t5))
}

/*
	算法复杂度分析

		Time complextiy : O(n)
		Space complexity : O(1)

	思路总结

		此题目用到一些数组的通用思路：

		1）题目要求in-place，就说明操作只能是原数组的元素的交换，或者元素的覆盖
		2）通过若干个额外的变量来记录一些状态（需要对题目进行分析）

	solution是否已读

		已读，solution和下面的实现是一致的。

	top 5 votes discuss 是否已读

		已读，吸收了top1的实现，见removeDuplicates_v1

*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		nums[idx] = nums[i]
		idx++
	}
	return idx
}

/*
	votes最多的discuss给出的实现，实现更为简洁一点

	记录重复的个数
	那么去重数组的下标就是 i - 重复的个数
*/
func removeDuplicates_v1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			continue
		}
		nums[i-count] = nums[i]
	}
	return len(nums) - count
}
