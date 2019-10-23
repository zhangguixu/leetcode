package main

import "fmt"

func main() {
	merge([]int{1,2,3,0,0,0,}, 3, []int{2,5,6,}, 3)
}

/*
	将数组num2合并到num1中

	nums1 ：[2，5，6, 0, 0, 0] 
	nums2： [3,4,7]

	在题目中，可以知道nums1有多余空间可以容纳nums2的元素。

	采用倒排的思路
*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := m - 1
	j := n - 1
	idx := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]
			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
		idx--
	}
	for j >= 0 {
		nums1[idx] = nums2[j]
		idx--
		j--
	}
	fmt.Println(nums1)
}