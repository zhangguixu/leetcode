package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 1, 0, 3, 12, 0, 0, 0})
	moveZeroes([]int{1, 1, 1, 3, 12, 0, 0, 0})
	moveZeroes_1([]int{0, 1, 0, 3, 12})
	moveZeroes_1([]int{0, 1, 0, 3, 12, 0, 0, 0})
	moveZeroes_1([]int{1, 1, 1, 3, 12, 0, 0, 0})
}

/*
	交换数值的解法
	runtime 64 ms
	momery: 7.7 MB
*/
func moveZeroes(nums []int) {

	l := len(nums)

	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			continue
		}
		allZero := true
		for j := i + 1; j < l; j++ {
			if nums[j] != 0 {
				allZero = false
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
		if allZero {
			break
		}
	}

	fmt.Println(nums)
}

/*
	将不是0的一一添加到前面，后续的补0即可，
	也就是两个index就可以解决了
	runtime 84ms
	momery: 8 MB
*/
func moveZeroes_1(nums []int) {
	lastFoundZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastFoundZero] = nums[i]
			lastFoundZero++
		}
	}
	for lastFoundZero < len(nums) {
		nums[lastFoundZero] = 0
		lastFoundZero++
	}
	fmt.Println(nums)
}
