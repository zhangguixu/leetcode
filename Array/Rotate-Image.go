package main

import "fmt"

func main() {
	rotate([][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	})
}

/*
	解题思路
	如果不允许申请额外的空间，必然是进行数组元素的交换，需要找到交换的规律，
	如何找到交换的规律？
	一个思路是 观察数组的旋转过程，观察90度的旋转
	[i][j] -> [n - j - 1][i]
	参考博文：https://www.cnblogs.com/grandyang/p/4389572.html
*/
func rotate(matrix [][]int)  {
	n := len(matrix)
  for i := 0; i < n/2 ; i++ {
		for j := i; j < n - i - 1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n - j - 1][i]
			matrix[n - j - 1][i] = matrix[n - i - 1][n - j - 1]
			matrix[n - i - 1][n - j - 1] = matrix[j][n - i -1]
			matrix[j][n - i - 1] = tmp
		}
	}
	fmt.Println(matrix)
}

// 借助额外空间
func rotateNotInPlace(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	columns := len(matrix[0])
	if rows != columns {
		return
	}
	m := make([][]int, 0, rows)
	for i := 0; i < rows; i++ {
		r := make([]int, 0, rows)
		for j := columns - 1; j >= 0; j-- {
			r = append(r, matrix[j][i])
		}
		m = append(m, r)
	}
	matrix = m
	fmt.Println(matrix)
}