package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}

/*
	题目要求，需要在线性时间内完成，即时间复杂度是O(n)

	思路是维护前三个最大的数，注意一下去重问题即可

*/
func thirdMax(nums []int) int {

	max, secMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64
	
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		switch {
		case n == max || n == secMax || n == thirdMax:
		case n > max:
			tmp := max
			max = n
			thirdMax = secMax
			secMax = tmp
		case n > secMax:
			thirdMax = secMax
			secMax = n
		case n > thirdMax:
			thirdMax = n
		}
	}

	if thirdMax > math.MinInt64 {
		return thirdMax
	}

	return max
}