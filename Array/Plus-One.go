package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{9, 9, 9}))

}

// 只要考虑到进位
func plusOne(digits []int) []int {

	total := len(digits)
	if total == 0 {
		return []int{1}
	}

	for total > 0 {
		total--
		digits[total] += 1
		if digits[total] < 10 || (digits[total] >= 10 && total == 0) {
			break
		}
		if digits[total] >= 10 {
			digits[total] -= 10
		}
	}

	if digits[0] >= 10 {
		digits[0] -= 10
		tmpList := make([]int, 0, len(digits)+1)
		tmpList = append(tmpList, 1)
		tmpList = append(tmpList, digits...)
		digits = tmpList
	}

	return digits
}
