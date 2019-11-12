package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
	fmt.Println(fourSum([]int{-1,0,1,2,-1,-4}, -1)) // [[-4 0 1 2] [-1 -1 0 1]]
	fmt.Println(fourSum([]int{1,-2,-5,-4,-3,3,3,5}, -11)) // [[-5 -4 -3 1]]
}

/*
	k数之和的套路

	先固定 k - 2个数（外循环，在外循环做一下判断进行去重和加速运算）
	然后剩下2个数，使用双指针法

	可以将算法的时间复杂度降低一点。
*/
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0, 100)
	sort.Ints(nums)
	for i := 0; i < len(nums) - 3; i++ {
		if nums[i] * 4 > target { // 判断条件，进行运算加速
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
	
		for j := i + 1; j < len(nums) - 2; j++ {
			if nums[i] + nums[j] * 3 > target { // 判断条件，进行运算加速
				break
			}
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l + 1] {
						l++
					}
					for l < r && nums[r] == nums[r - 1] {
						r--
					}
					l++
					r--
					continue
				}
				if sum > target {
					r--
					continue
				}
				l++
			}
		}
	}
	return res
}