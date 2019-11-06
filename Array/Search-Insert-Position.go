package main

import "fmt"

func main() {
	fmt.Println(searchInsert_v1([]int{1,3,5,6}, 2)) // 1
	fmt.Println(searchInsert_v1([]int{1,3,5,6}, 0)) // 0
	fmt.Println(searchInsert_v1([]int{1,3,5,6}, 7)) // 4
}

/*
	给定的数组是有序的，当然可以遍历nums数组，比较之后获取插入数据下标

	最差的情况是遍历完整个数组

	算法的时间复杂度是O(n)

	注意考虑边界
	1） 当输入的值是最小的时候
	2） 当输入的值是最大的时候
*/
func searchInsert(nums []int, target int) int {
	var i int
  for i = 0; i < len(nums); i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
	}
	return i
}

/*
	算法的时间复杂度可以降低到O(logn)

	利用二叉查找的思想，缩小查找的范围

	这里的边界问题比较难处理，不然就得增加if else

	1）r的初始值
	2）循环结束的判断
	3）l的赋值
*/
func searchInsert_v1(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return r
}