package main

import (
	"fmt"
	"math"
)

func main() {
	var n1 []int
	n1 = []int{4,3,2,1}
	nextPermutation(n1)
	fmt.Println(n1) // [1 2 3 4]

	n1 = []int{4,2,1,3}
	nextPermutation(n1)
	fmt.Println(n1) // [4 2 3 1]

	n1 = []int{4,1,3,2}
	nextPermutation(n1)
	fmt.Println(n1) // [4 2 1 3]

	n1 = []int{4,2,0,2,3,2,0}
	nextPermutation(n1)
	fmt.Println(n1) // [4,2,0,3,0,2,2]

	n1 = []int{5,4,7,5,3,2} 
	nextPermutation(n1)
	fmt.Println(n1) // [5,5,2,3,4,7]
}

/*
	题目要求inplace，只允许空间复杂度为O(n)，那么肯定是指针法 

	先列举尽可能多例子，观察规律

	4,3,2,1 => 1,2,3,4 

	4,2,1,3 => 4,2,3,1

	4,1,3,2 => 4,2,1,3 

	4,2,3,1 => 4,3,1,2

	算法步骤说明：

	从低位开始（也就是倒序），找到第一个小于右侧数字的元素，

	这个就是需要交换的数字，找到这个元素右侧大于这个元素的最小值，进行交换

	然后右侧的元素再进行升序
*/
func nextPermutation_origin(nums []int) {
	idx := -1
	n := len(nums) - 1
	for i := n; i > 0; i-- {
		if nums[i] > nums[i - 1] {
			idx = i - 1
			break
		}
	}
	if idx == -1 {
		nums[0], nums[n] = nums[n], nums[0]
		idx++
	} else {
		min := math.MaxInt64
		minIdx := -1
		for i := n; i > idx; i-- {
			if nums[i] > nums[idx] && min > nums[i] {
				min = nums[i]
				minIdx = i
			}
		}
		nums[idx], nums[minIdx] = nums[minIdx], nums[idx]
	}
	for i := n; i > idx; i-- {
		for j := i - 1; j > idx; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

/*
	这个算法步骤可以改进的

	因为是倒序一路向上找元素小于右侧数字，因此找到的元素的右侧其实是一个降序序列！！！！！ 优化的核心点

	这道题考察在于对数组序列的敏感度
*/
func nextPermutation(nums []int) {
	var i int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i + 1] {
			break
		}
	}
	if i >= 0 {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	reverse(nums, i + 1, len(nums) - 1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}