package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{1, 1}, []int{1,2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	min := len(nums1)
	if min < len(nums2) {
		min = len(nums2)
		nums1, nums2 = nums2, nums1
	}
	numMap := make(map[int]int, min)
  list := make([]int, 0, min)
	for i := 0; i < min; i++ {
		numMap[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		count := numMap[nums2[i]]
		if count > 0 {
			numMap[nums2[i]]--
			list = append(list, nums2[i])
		}
	}
	return list
}