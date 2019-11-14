package main

import "fmt"

func main() {
	// fmt.Println(search([]int{4,5,6,7,0,1,2,3}, 0)) // 4
	// fmt.Println(search([]int{4,5,6,7,0,1,2}, 3)) // -1
	// fmt.Println(search([]int{4,5,6,7,0,1,2}, 0)) // 4

	fmt.Println(search_v1([]int{4,5,6,7,0,1,2,3}, 0)) // 4
	// fmt.Println(search([]int{4,5,6,7,0,1,2}, 3)) // -1
	// fmt.Println(search([]int{4,5,6,7,0,1,2}, 0)) // 4

	// fmt.Println(searchPivot([]int{4,5,6,7,0,1,2,3})) // 4
	// fmt.Println(searchPivot([]int{4,5,6,7,8,0,1,2})) // 5
	// fmt.Println(searchPivot([]int{0,1,2,3,4,5,6,7})) // 5


	// fmt.Println(binarySearch([]int{0,1,2,3,4,5,6}, 3))
	// fmt.Println(binarySearch([]int{0,1,2,3,4,5,6}, 8))
	// fmt.Println(binarySearch([]int{0,1,2,3,4,5,6}, -1))
}

/*
	一个降序的数组被旋转，旋转点不知

	这里需要查找一个数字，要求时间复杂度为O(logN)

	对于一个选择的降序数组

	[4,5,6,7,0,1,2,3]

	如果nums[len(nums) - 1] > nums[0] 说明数组没有旋转（或者刚好旋转一圈）

	如果target > nums[0] 的话，说明target是在数组的前半部分

	如果target < nums[0] 的话，说明target是在数组的后半部分

	思路是
	1）使用类似于二分查找的方法找到旋转点
	2）然后在指定的范围，再使用二分查找
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	pivot := searchPivot(nums)
	if pivot == len(nums) && (target < nums[0] || target > nums[pivot - 1]) {
		return -1
	}

	var l, r int

	if target >= nums[0] {
		l = 0
		r = pivot - 1
	} else {
		l = pivot
		r = len(nums) - 1
	}

	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		}
		if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

/*
	寻找旋转点
	类似于二分查找，但是有点不同
	这里还需要多对比几种情况
*/
func searchPivot(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[l] < nums[m] {
			l = m
		} else {
			r = m
		}
	}
	return l + 1
}

/*
	二分查找
*/
func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		fmt.Println(l, m, r)
		if target == nums[m] {
			return m
		}
		if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

/*
	在题解中，找到一个最简洁的解法

	思路是进行一次二分查找

	nums[0] <= nums[mid]（0 - mid不包含旋转）且nums[0] <= target <= nums[mid] 时 high 向前规约；

	nums[mid] < nums[0]（0 - mid包含旋转），target <= nums[mid] < nums[0] 时向前规约（target 在旋转位置到 mid 之间）

	nums[mid] < nums[0]，nums[mid] < nums[0] <= target 时向前规约（target 在 0 到旋转位置之间）

	其他情况向后规约

	也就是说nums[mid] < nums[0]，nums[0] > target，target > nums[mid] 三项均为真或者只有一项为真时向后规约。

	作者：LukeLee
	链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/ji-jian-solution-by-lukelee/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func search_v1(nums []int, target int) int {
	// TBD
}
