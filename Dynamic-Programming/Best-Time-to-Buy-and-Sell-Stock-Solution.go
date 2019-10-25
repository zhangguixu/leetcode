package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit_v2([]int{7,1,5,3,6,4,}))
	fmt.Println(maxProfit_v2([]int{7,6,4,3,1,}))
	fmt.Println(maxProfit_v2([]int{7,2,4,1,6,}))
}

/*
	动态规划(dynamic programming)是运筹学的一个分支，
	是求解决策过程(decision process)最优化的数学方法。
	
	20世纪50年代初美国数学家R.E.Bellman等人在研究多阶段决策过程(multistep decision process)的优化问题时，
	提出了著名的最优化原理(principle of optimality)，把多阶段过程转化为一系列单阶段问题，
	利用各阶段之间的关系，逐个求解，创立了解决这类过程优化问题的新方法——动态规划

	资料：https://cloud.tencent.com/developer/article/1086657


	说实话，在思考这个题目时候，感觉不是典型的动态规划题目，不符合整体动态规划的解决常规步骤，这里仅仅只是涉及到了多阶段决策，并不能拆分成多个子问题。
	存在误导的情况，导致问题容易想歪。

	在给定的股票价格中，需要决定买不买，买了之后，需要觉得什么时候卖，如果按照每天去观察计算的话，可能性太多了，复杂度很高。

	那么从整体来看，只允许买卖一次，那么就从最低点买入，从最高点卖出即可，这里就转换为两个状态

	* 最低价格 minPrice
	* 最高价格 maxPrice 

	因为结果是求最高收益，那么最高价格可以转化为 maxPrice minPrice + maxProfit 

	而且必须先买，再卖，这里就是一次循环（i不能后退）

*/
func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i] - minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

/*
	上面说的这道题如果用动态规划存在误导，只是当时找不到推导的公式

	假设dp[i]是第i天卖出股票的最大利润，那就是

	dp[i] = max(arr[i] - arr[i - 1] + dp[i - 1] , 0)
	dp[0] = 0

	题目要求的就是最大利润，因此就是

	max(dp[0], ..., dp[i])

	获得递推式之后，可以开始写代码了
*/
func maxProfit_v2(prices []int) int {
	maxPro := 0
	curMax := 0
	if len(prices) == 1 {
		return maxPro
	}
	for i := 1; i < len(prices); i++ {
		curMax = prices[i] - prices[i - 1] + curMax
		if curMax < 0 {
			curMax = 0
		}
		if maxPro < curMax {
			maxPro = curMax
		}
	}
	return maxPro
}