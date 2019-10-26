package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1,2,3,1,}))
	fmt.Println(rob([]int{2,7,9,3,1,}))
	fmt.Println(rob([]int{2,1,1,2,}))
	fmt.Println(rob([]int{1,2,3,1,}))

	fmt.Println("==============")

	fmt.Println(rob_v2([]int{1,2,3,1,}))
	fmt.Println(rob_v2([]int{2,7,9,3,1,}))
	fmt.Println(rob_v2([]int{2,1,1,2,}))
	fmt.Println(rob_v2([]int{1,2,3,1,}))

	fmt.Println("==============")

	fmt.Println(rob_v3([]int{1,2,3,1,}))
	fmt.Println(rob_v3([]int{2,7,9,3,1,}))
	fmt.Println(rob_v3([]int{2,1,1,2,}))
	fmt.Println(rob_v3([]int{1,2,3,1,}))
}

/*
	使用动态规划进行解答，首先分为两个步骤

	1） 当第i个被抢的情况下，可以获得最大收益的推导公式

		假设DP[i]表示第i个被抢的最大收益，根据题目的规则，不能抢劫连续的房子，有以下公式：

		DP[i] = max(DP[i - 2], DP[i - 3], DP[0]) + arr[i] 

		DP[0] = arr[0]

		DP[1] = arr[1]

	2） 获取的最大收益是max(DP[0], ..., DP[i])

*/
func rob(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		if i <= 1 {
			dp[i] = nums[i]
		} else {
			maxDp := 0
			for j := i - 2; j >= 0; j-- {
				if maxDp < dp[j] {
					maxDp = dp[j]
				}
			}
			dp[i] = maxDp + nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

/*
	以上关于maxDp操作可以简化
*/
func rob_v2(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	max := 0
	maxDp := 0
	for i := 0; i < len(nums); i++ {
		if i <= 1 {
			dp[i] = nums[i]
		} else {
			if maxDp < dp[i - 2] {
				maxDp = dp[i - 2]
			}
			dp[i] = maxDp + nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

/*
	dp数组其实每次用到就两位，可以减少dp数组占用的空间
*/
func rob_v3(nums []int) int {
	dp := make([]int, 2, 2)
	maxDp := 0
	max := 0
	cur := 0
	for i := 0; i < len(nums); i++ {
		if i <= 1 {
			cur = nums[i]
			dp[i] = nums[i]
		} else {
			if maxDp < dp[0] {
				maxDp = dp[0]
			}
			cur = maxDp + nums[i]
			dp[0] = dp[1]
			dp[1] = cur
		}
		if max < cur {
			max = cur
		}
	}
	return max
}



