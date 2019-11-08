package main

import "fmt"

func print(nums []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf(" %d ", nums[i])
	}
	fmt.Printf("\n")
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	print(nums, removeElement(nums, 2))
}

/*
	two pointer
*/
func removeElement(nums []int, val int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
