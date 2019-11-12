package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, -1, 0, 1, 1}))
	fmt.Println(threeSum([]int{0,0,0}))
	fmt.Println(threeSum([]int{-2,1,1}))
	fmt.Println(threeSum([]int{-2,0,1,1,2}))
}

/*
	这道题

	1）利用排序来降低难度
	2）如何处理三个在变化的数，当面对三个变数时，先确定其中的一个（循环），然后再用双指针，
	3）在1），2）的基础上处理去重就简单很多

	这道题跟easy的不太一样的地方，需要多个技巧一起运用，而且逻辑边界很多，如果思路不正确的话，很容易陷入无限的修复边界情况
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 100)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ { // 完成的关键
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
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
			if sum > 0 {
				r--
				continue
			}
			l++
		}
	}
	return res
}


