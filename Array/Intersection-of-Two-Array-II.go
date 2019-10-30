package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intersect_v1([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect_v1([]int{1, 1}, []int{1, 2}))
}

/*
	hashtable：记录元素出现的次数

	时间复杂度： O(n)，需要遍历给定的两个数组
	空间复杂度： O(n)，需要存储数组长度较小的元素

	这里需要注意的元素可能是重复出现的，因此hashtable是要对元素进行计数的。
*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	countMap := make(map[int]int, len(nums1))
	result := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		countMap[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		n := nums2[i]
		if countMap[n] > 0 {
			countMap[n]--
			result = append(result, n)
		}
	}
	return result
}

/*
	这道题的solution是需要收费查看的，穷就算了。

	根据题目的follow up，如果给定的数组是已经排好序的，那么要如何处理？

	这个时候就不用额外的空间了

	时间复杂度： O(n) ~ O(nlogn)
	空间复杂度： O(1)
*/
type List []int

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	return l[i] < l[j]
}

func intersect_v1(nums1 []int, nums2 []int) []int {
	sort.Sort(List(nums1))
	sort.Sort(List(nums2))
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	result := make([]int, 0, len(nums1))
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			j++
			continue
		}
		result = append(result, nums1[i])
		i++
		j++
	}
	return result
}