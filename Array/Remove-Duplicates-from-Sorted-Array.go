package main

import "fmt"

func main() {
	removeDuplicates([]int{1, 1, 2})
	removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
	removeDuplicates([]int{1,1,1}) // 边界
	removeDuplicates([]int{1})	// 边界
	removeDuplicates([]int{}) // 边界
}

func removeDuplicates(nums []int) int {
	
	var i, j int
	total := len(nums)

	if total <= 1 {
		return total
	}

	for i < total - 1 {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
		i++ 
	}

	if j > 0 && nums[j - 1] != nums[i] {
		nums[j] = nums[i]
		j++
	}

	if j == 0 {
		j = 1
	}


	fmt.Println(nums, j)

	return j
}