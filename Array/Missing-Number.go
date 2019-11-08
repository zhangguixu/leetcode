package main 

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3,0,1}))
	fmt.Println(missingNumber([]int{3,0,1}))
	fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1}))
}

/*
	题目要求

	1）时间复杂度为O(n)
	2）空间复杂度为O(1)

	利用差即可
*/
func missingNumber(nums []int) int {
	var actualSum, totalSum int
	for i := 0; i < len(nums); i++ {
		actualSum += nums[i]
		totalSum += i
	}
	totalSum += len(nums)
	return totalSum - actualSum
}