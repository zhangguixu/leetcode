package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	var max, count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
			continue
		}
		count++
		if max < count {
			max = count
		}
	}
	return max
}