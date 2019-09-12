package main

import "fmt"

func main() {
	rotate_1([]int{1, 2, 3, 4, 5, 6, 7}, 1)
	rotate_1([]int{1, 2, 3, 4, 5, 6, 7}, 2)
	rotate_1([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	rotate_1([]int{1, 2}, 3)
	rotate_1([]int{1, 2, 3}, 4)
}

// 题目要求 nums的地址不能变，因此，不能对nums进行重新赋值，也不能对nums进行的append
// 因此，使用copy将第二个slice里的元素拷贝到第一个slice里，拷贝的长度为两个slice中长度较小的长度值
// 由于可以一直旋转，因此，这边实际上是取摸的操作
// runtime: 72ms Mem: 7.6 MB
func rotate(nums []int, k int) {
	n := len(nums)
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[(i+k)%n] = nums[i]
	}
	copy(nums, list)

	fmt.Println(nums)
}

// rumtine: 52 ms Mem: 7.7MB
func rotate_1(nums []int, k int) {
	l := len(nums) - k%len(nums)
	list := nums[l:]
	for i := 0; i < l; i++ {
		list = append(list, nums[i])
	}
	copy(nums, list)

	fmt.Println(nums)
}
