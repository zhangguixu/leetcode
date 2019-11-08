package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{9, 9, 9}))

}

/*
	按照题目的要求，就是使用数组来表示数字，然后执行+1操作，
	需要考虑的就是进位的问题
*/

/*
	数组的倒序访问:

	这里的复杂度波动比较大

	时间复杂度: O(n)
	空间复杂度: O(n)
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] <= 9 {
			return digits
		}
		digits[i] = 0
	}
	tmpList := make([]int, 0, len(digits)+1)
	tmpList = append(tmpList, 1)
	tmpList = append(tmpList, digits...)
	digits = tmpList
	return digits
}

/*
	进行反转，正序访问

	这里有一个好处就是slice的append操作不一定需要申请额外的空间

	时间复杂度为：O(n)
	空间复杂度：O(1)
*/
func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func plusOne_v1(digits []int) []int {
	n := len(digits)
	reverse(digits)
	digits[0] += 1
	for i := 1; i < n; i++ {
		if digits[i-1] <= 9 {
			break
		}
		digits[i-1] -= 10
		digits[i] += 1
	}
	if digits[n-1] >= 10 {
		digits[n-1] -= 10
		digits = append(digits, 1)
	}
	reverse(digits)
	return digits
}

/*
	在discuss找到了改进第一个做法的例子，

	在数组 + 1的操作中，需要进位的情况其实只有

	9 + 1 = 10，那么元素的值就是0，进一位

	因此如果是整体都进位时，只需要在数组末端append一个0
*/
func plusOne_v2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] <= 9 {
			return digits
		}
		digits[i] = 0
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
