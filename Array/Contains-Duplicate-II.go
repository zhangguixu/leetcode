package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1,0,1,1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
}

/*
	这套题的题目描述有点坑，导致ac率有点低

	给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，且 |i - j| <= k

	nums = [1,2,3,1], k = 3

	hashtable

	时间复杂度为O(n)
	空间复杂度为o(n)
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	countMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		pre, ok := countMap[n]
		if ok && i - pre <= k {
			return true
		}
		countMap[n] = i
	}
	return false
}