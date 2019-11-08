package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
	fmt.Println(maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
	fmt.Println(maxProfit([]int{2, 4, 1}))
}

/*
	算法复杂度分析

		Time complextiy : O(n)
		Space complexity : O(1)

	思路总结

		*此题的关键在于画图，从图表的变化可以很快找到答案

		Peak Valley Approach，找到每一个起伏中的波峰和波谷，在波谷进行买入，在波峰卖出，获取局部的最大利润，

		将局部最大利润累加，获取全部最大的利润

	solution是否已读

		solution给出了三种方法

		1） brute Force 暴力破解

		2） Peak Valley Approach

		也是本人解法思路，在实现上，可以更加简洁，
		见maxProfit_v1，原solution给的解法是在数组为空的时候，会runtime error

		3） Peak Valley Approach  改进版

		见 maxProfit_v2

	top 5 votes discuss 是否已读

		已读，都是solution中提到的解法

*/
func maxProfit(prices []int) int {
	total := len(prices)
	profit := 0
	for i := 0; i < total-1; i++ {
		if prices[i] > prices[i+1] {
			continue
		}
		for j := i + 1; j < total; j++ {
			if j < total-1 && prices[j] < prices[j+1] {
				continue
			}
			if j == total-1 && prices[j] < prices[j-1] {
				continue
			}
			profit += prices[j] - prices[i]
			i = j
			break
		}
	}
	return profit
}

func maxProfit_v1(prices []int) int {
	maxprofit := 0
	valley := 0
	peak := 0
	i := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxprofit += peak - valley
	}
	return maxprofit
}

/*
	依旧是Peak Valley Approach，

	波峰 - 波谷 = 每一小段的差值累加
*/
func maxProfit_v2(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}
