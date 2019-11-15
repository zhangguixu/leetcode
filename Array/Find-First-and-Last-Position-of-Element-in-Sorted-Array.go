package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))

	fmt.Println("=======================================")

	fmt.Println(searchRange_v1([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange_v1([]int{5,7,7,8,8,10}, 6))
}

/*
	二分查找，找到对应的值之后，向前向后探索范围（线性）
*/
func searchRange(nums []int, target int) []int {
	l, m, r := 0, 0, len(nums) - 1
	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			i, j := m, m
			for i < len(nums) && nums[i] == target {
				i++
			}
			for j >= 0 && nums[j] == target {
				j--
			}
			return []int{j + 1, i - 1}
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return []int{-1, -1}
}

/*
	上述的实现有一个线性查找的，可能最差的情况，算法的时间复杂度为O(n)，例如[4,4,4,4,4,4]
	因此这边使用二分查找到范围的最左最右的
*/
func searchRange_v1(nums []int, target int) []int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums) - 1] {
		return []int{-1, -1}
	}
	return []int{findEdgeIndex(nums, target, true), findEdgeIndex(nums, target, false)}
}

func findEdgeIndex(nums []int, target int, left bool) int {
	l, m, r, edge := 0, 0, len(nums) - 1, -1
	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			edge = m
			if left {
				r = m - 1
			} else {
				l = m + 1
			}
			continue
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return edge
}