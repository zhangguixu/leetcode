package main

import "fmt"

func main() {

}


/*
	难点还是在于如何处理去重的问题

	使用hashtable进行加速

*/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0, 100)
	if len(candidates) == 0 || target < candidates[0] {
		return res
	}
	numMap := make(map[int]int, len(candidates))
	for i := 0; i < len(candidates); i++ {
		numMap[candidates[i]] = i
	}
	for i := 0; i < len(candidates); i++ {
		n := candidates[i]
		divisor := target / n
		for divisor >= 1 {
			left := target - n * divisor
			if left == 0 {
				list := make([]int, 0,  divisor)
				for j := 0; j < divisor; j++ {
					list = append(list, n)
				}
				continue
			}
		
			divisor--
		}
	}
}