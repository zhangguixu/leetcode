package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4,}))
	// fmt.Println(maxSubArray([]int{8, -19, 5, -4, 20}))
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4,}))
}

/*
	O(n)的算法，称为Kadane’s Algorithm，下面的文章对其分析
	https://hackernoon.com/kadanes-algorithm-explained-50316f4fd8a6

	思路的起点是 假设
	dp[i] 是以i为子数组终点的最大值，那么就有以下公式

	dp[i] = max(dp[i - 1], 0) + arr[i]

	dp[0] = arr[0]

	题目最后所求的 max(dp[0], ... , dp[i]) 
*/

/*
	递归的实现，可以看到很多重叠子问题的求解，导致整体runtime异常的高，
	在leetcode提交了之后的runtime为612ms
*/
func maxSubArray(nums []int) int {
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		m := helper(nums, i)
		if max < m {
			max = m
		}
	}
	return max
}

func helper(nums []int, n int) int {
	if n == 0 {
		return nums[0]
	}
	preMax := helper(nums, n - 1)
	// fmt.Println(preMax, n)
	return max(preMax, 0) + nums[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


/*
	以上的解法是自顶向下的，可以使用自低向上递推来做

	我们需要记录的是
	* 以当前下标为终点的子数组最大值 cur
	* 所有子数组的最大值 max （题目所求）
*/
func maxSubArray1(nums []int) int {
	var cur int
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if max < cur {
			max = cur
		}
		if cur < 0 {
			cur = 0
		}
	}
	return max
}
