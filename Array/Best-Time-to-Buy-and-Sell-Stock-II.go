package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4,}))
	fmt.Println(maxProfit([]int{7,6,4,3,1,}))
	fmt.Println(maxProfit([]int{1,2,3,4,5,}))
	fmt.Println(maxProfit([]int{3,2,6,5,0,3,}))
	fmt.Println(maxProfit([]int{2,1,2,1,0,1,2}))
}

/*
	贪心算法:
	追求局部最优
	1.建立数学模型来描述问题。
    2.把求解的问题分成若干个子问题。
    3.对每一子问题求解，得到子问题的局部最优解。
	4.把子问题的解局部最优解合成原来解问题的一个解。
	
	贪心算法的证明围绕着：整个问题的最优解一定由在贪心策略中存在的子问题的最优解得来的

	买卖股票的策略

	将数组中的值展开为一个曲线图，可以看到整体一个发展趋势，波峰 - 波低
*/
// func maxProfit(prices []int) int {
// 	maxP := 0
// 	total := len(prices)
// 	for i := 0; i < total; i++ {
// 		profit := 0
// 		for j := i; j < total; j++ {
// 			if j+1 < total && prices[j] > prices[j+1] {
// 				continue
// 			}
// 			crestIndex := findCrest(prices, j, total - 1)
// 			if crestIndex == -1 {
// 				continue
// 			}
// 			fmt.Println(prices[crestIndex], prices[j])
// 			profit += prices[crestIndex] - prices[j]
// 			j = crestIndex
// 		}
// 		if profit > maxP {
// 			maxP = profit
// 		}
// 	}
// 	return maxP
// }

// func findCrest(prices []int, start, end int) (crestIndex int) {
// 	crestIndex = -1
// 	i := start + 1
// 	for i < end {
// 		if prices[start] < prices[i] && prices[i] > prices[i+1] {
// 			crestIndex = i
// 			break
// 		}
// 		i++
// 	}
// 	if i == end && prices[i] > prices[start] {
// 		fmt.Println("--", i, start)
// 		crestIndex = i
// 	}
// 	return crestIndex
// }
func maxProfit(prices []int) int {
	total := len(prices)
	profit := 0
	for i := 0; i < total - 1; i++ {
		if prices[i] > prices[i+1] {
			continue
		}
		for j := i+1; j < total; j++ {
			if j < total - 1 && prices[j] < prices[j+1] {
				continue
			}
			if j == total -1 && prices[j] < prices[j - 1] {
				continue
			}
			profit += prices[j] - prices[i]
			i = j
			break
		}
	}
	return profit
}
