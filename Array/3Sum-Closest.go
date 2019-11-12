package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 0))
	fmt.Println(threeSumClosest([]int{1,1,1,1}, -100))
	fmt.Println(threeSumClosest([]int{-55,-24,-18,-11,-7,-3,4,5,6,9,11,23,33}, 0))
}

/*
	跟3sum一样的思路和解法，只是判断的条件有所调整
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var closest, close, sum, l, r int
	for i := 0; i < len(nums) - 2; i++ {
		l = i + 1
		r = len(nums) - 1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if l == i + 1 && r == len(nums) - 1 {
				close = sum
			}
			if abs(close, target) > abs(sum, target) {
				close = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
		if i == 0 {
			closest = close
		}
		if abs(closest, target) > abs(close, target) {
			closest = close
		}
	}
	return closest
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

