package main

import "fmt"

func main() {
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}

/*
	题目一开始要求就是in-place，那么要实现数组的旋转，就应该是数组元素的交换实现，
	这也说明数组的旋转中，元素的交换应该是有规律，这里需要的是找出规律。

	先从这个例子进行观察

	[
  	[1,2,3],
  	[4,5,6],
  	[7,8,9]
	]

	旋转之后

	[
  	[7,4,1],
  	[8,5,2],
  	[9,6,3]
	]

	观察元素1，3，9，7，数组旋转，这四个元素刚好是互换位置，其坐标的变化是

	1 [0,0] -> 3 [0,2] -> 9 [2,2] -> 7 [2,0] -> 1 [0,0]

	假设 i = 0, j = 0 那么就有规律

	1 [i, j] -> 3 [j , n - i - 1] -> 9 [n - i - 1, n - j - 1] -> 7 [n - j - 1, n - (n - i - 1) - 1 = i] -> [i, n - (n - j - 1) - 1 = j]

	这里找规律的过程耗时比较长，主要是靠猜测和运算验证。

	找到坐标变化的规律之后，发现这里是一个四值交换位置的问题。

	接下来就是确定i和j的取值，确保已经旋转的元素不再进行交换。

	i < n/2

	i <= j < n - i - 1  这个值是在交换的过程中不断尝试得来的

	https://www.youtube.com/watch?v=Jtu6dJ0Cb94  这个视频非常详细的解释了这个实现的思考过程，可以看看

*/
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = tmp
		}
	}
	fmt.Println(matrix)
}

/*
	discuss中投票最多的答案，给出了两种旋转的方式，但是有一个问题是：这是如何想到的。

	1） 顺时针旋转：first reverse up to down, then swap the symmetry

	1 2 3     7 8 9     7 4 1
 	4 5 6  => 4 5 6  => 8 5 2
	7 8 9     1 2 3     9 6 3

	2） 逆时针旋转：first reverse left to right, then swap the symmetry

	1 2 3     3 2 1     3 6 9
 	4 5 6  => 6 5 4  => 2 5 8
 	7 8 9     9 8 7     1 4 7

*/
func rotate_v1(matrix [][]int) {
	reverse(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverse(matrix [][]int) {
	i := 0
	j := len(matrix) - 1
	for i < j {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
}
