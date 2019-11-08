package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
	fmt.Println(findDisappearedNumbers_v1([]int{4,3,2,7,8,2,3,1}))
}

/*
	数组的元素 1 <= a[i] <= n

	n是数组的长度

	找出缺失的元素

	要求
	空间复杂度: O(1)
	时间复杂度: O(n)

	思路，
	数组的元素和下标存在关系：如果没有缺失的情况下，将数组元素的值 - 1 作为下标可以有且只有一次访问到数组的所有元素，即

	for i in [0, n - 1]

	a[a[i] - 1] 可以遍历到数组的所有元素

	举个例子 [3,1,2]

	i        a[i]         a[a[i] - 1]
	0         3               2
	1         1               3
	2         2               1

	如果存在重复的情况下，例如 [3,1,1]

	i        a[i]         a[a[i] - 1]
	0         3               2
	1         1               3
	2         1               3

	a[i] - 1代表的下标会被重复访问，可以利用数组自身来记录元素是否已经通过a[i] - 1的方式访问过了

	这里采用的方式是 将a[a[i] - 1] =  a[a[i] - 1] * (n + 1) 来标记已经访问过的状态
*/
func findDisappearedNumbers(nums []int) []int {
	total := len(nums)
	res := make([]int, 0)
	for i := 0; i < total; i++ {
		idx := nums[i]
		if idx > total {
			idx = idx / (total + 1)
		}
		if nums[idx - 1] <= total {
			nums[idx - 1] = nums[idx - 1] * (total + 1)
		}
	}
	for i := 0; i < total; i++ {
		if nums[i] <= total {
			res = append(res, i + 1)
		}
	}
	return res
}

/*
	在讨论中看到另一种解法，核心思路跟上述的一样的，就是数组的元素和下标之间存在着关系

	这里使用的是交换，就是将a[i] 元素放在 a[a[i] - 1]上，

	二次遍历的时候，如果a[i] != i + 1，说明就是要找的元素

	举个例子: [3,1,1,4]

	i        a[i]         a[i] <=> a[a[i] - 1]     a
	0         3            3 <=> 1                 [1,1,3,4]
	1         1            1 <=> 1                 [1,1,3,4]
	2         1            3 <=> 3                 [1,1,3,4]
	3         4            4 <=> 4                 [1,1,3,4]   

	还未完全理解，需要找个时间再看看
*/
func findDisappearedNumbers_v1(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i + 1 {
			res = append(res, i + 1)
		}
	}
	return res
}
