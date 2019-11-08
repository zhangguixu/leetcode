package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3, 3}, 6))
}

/*
	two pass hashtable
	借助额外的空间来减少时间复杂度
	原来也是借助hashtable实现，但是实现的有问题，特别的丑陋，其实是可以在原来的代码基础上优化的
	后续练习要多注意
*/
func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		idx, ok := m[complement]
		if ok && idx != i {
			res = []int{i, idx}
			break
		}
	}
	return res
}

/*
	One pass hashtable

	对上述方法的优化，减少一次遍历

	优化的思路：遍历的次数是否可以减少
*/
func twoSum_v2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var res []int
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		idx, ok := m[complement]
		if ok {
			res = []int{idx, i}
			break
		}
		m[nums[i]] = i
	}
	return res
}

/*
	数组解题技巧： 有序数组是否可以降低解决问题的难度

	对于一个有序的数组arr, i = 0, arr[len(arr) - 1]

	若 arr[i] + arr[j] > target 那么 j--, i 不变
	若 arr[i] + arr[j] < target 那么 i++, j 不变
	直到 arr[i] + arr[j] == target 为止

	这道题有一个麻烦的点，那就是返回的是数组的下标，我们对原数组进行排序，这个时候下标跟原来的不一致了。

	这里就需要额外的空间，来保存排序之后数组，找到数组元素之后，还需要遍历原数组，找到元素对应的下标，再返回。

	时间复杂度度：O(nlogN)
	空间复杂度：O(n)

	比不上使用hashtable的方便，

	此想法落地太丑陋了，忽略
*/
