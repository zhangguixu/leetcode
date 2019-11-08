package main

import "fmt"

func main() {
	fmt.Println(generate_v1(5))
}

/*
	杨辉三角计算
	numRows = 5

	[
		[1],
		[1,1],
		[1,2,1],
		[1,3,3,1],
		[1,4,6,4,1]
	]

	规律是：

	nums[i][j] = nums[i - 1][j - 1] + nums[i - 1][j]

	若 i - 1 < 0， 则nums[i][j] = 1, 即nums[0][0] = 1
	若 j - 1 < 0, nums[i][j] = nums[i - 1][j]
	若 j < i，则 nums[i][j] = nums[i - 1][j - 1]

	nums[0][0] = 1
*/

func generate(numRows int) [][]int {
	nums := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		nums[i] = make([]int, i+1, i+1)
		for j := 0; j <= i; j++ {
			if i-1 < 0 {
				nums[i][j] = 1
			}
			if j-1 >= 0 {
				nums[i][j] += nums[i-1][j-1]
			}
			if j < i {
				nums[i][j] += nums[i-1][j]
			}
		}
	}
	return nums
}

/*
	实现可以简洁一点
*/
func generate_v1(numRows int) [][]int {
	nums := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		nums[i] = make([]int, i+1, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				nums[i][j] = 1
			} else {
				nums[i][j] = nums[i-1][j-1] + nums[i-1][j-1]
			}
		}
	}
	return nums
}
