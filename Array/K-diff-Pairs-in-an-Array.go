package main

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2)) // 2
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0)) // 1
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1)) // 4
	fmt.Println(findPairs([]int{1,1,1,1,1}, 0)) // 1

	fmt.Println("===================================")

	// fmt.Println(findPairs_v1([]int{3, 1, 4, 1, 5}, 2)) // 2
	// fmt.Println(findPairs_v1([]int{1, 3, 1, 5, 4}, 0)) // 1
	// fmt.Println(findPairs_v1([]int{1, 2, 3, 4, 5}, 1)) // 4
	// fmt.Println(findPairs_v1([]int{1,1,1,1,1}, 0)) // 1
}

/*
	思路：

	先排序，再用指针法

	空间复杂度: O(1)
	时间复杂度：O(n*n) 
*/
func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	pre := math.MinInt64
	var cur, tail int
	for cur < len(nums) && tail < len(nums) {
		if pre == nums[cur] {
			cur++
			continue
		}
		if tail <= cur {
			tail = cur + 1
		}
		for j := tail; j < len(nums); j++ {
			if nums[cur] + k == nums[j] {
				tail = j + 1
				pre = nums[cur]
				count++
				break
			}
		}
		cur++
	}
	return count
}