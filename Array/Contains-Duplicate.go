package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsDuplicate_v1([]int{1, 2, 3, 1}))
}

/*
	算法复杂度分析

		时间复杂度是 O(n)
		空间复杂度是 O(n)

	使用hashtable进行元素的计数
*/
func containsDuplicate(nums []int) bool {
	numMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]]++
		if numMap[nums[i]] > 1 {
			return true
		}
	}
	return false
}

/*
	在solution中，除了暴力破解，还有一种解法是先给数组排序，然后在遍历数组，这个时候就只需要O(1)的空间，就可以找出是否有重复的元素

	这种实现的时间复杂度只要看排序的时间复杂度

	估算的时间复杂度： O(nlogN) ~ O(n)
	空间复杂度： O(1)

	Go语言中，对数组进行排序，需要实现Sort接口，sort中会自动选择排序算法，https://www.jb51.net/article/127463.htm
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

func containsDuplicate_v1(nums []int) bool {
	sort.Sort(List(nums))
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
