package main

import "fmt"

func main() {
	fmt.Println(getRow(5))
	fmt.Println(getRow_v1(5))
}

/*
	题目的follow要求是否可以将空间的复杂度降低为O(k)

	因为此题目求的是某一行的数组，而某一行的数组其实只跟它上一行的数组的数据有关

	还是同样的计算规律，只是不再使用一个二维数据进行存储每一行的数据
*/
func getRow(rowIndex int) []int {
	var curRow, lastRow []int
	for i := 0; i <= rowIndex; i++ {
		curRow = make([]int, 0, i + 1)
		for j := 0; j <= i; j++ {
			target := 1
			if j != 0 && j != i {
				target = lastRow[j - 1] + lastRow[j]
			}
			curRow = append(curRow, target)
		}
		lastRow = curRow[0:]
	}
	return curRow
}

/*
	上述实现还存在优化空间：

	可以将lastRow的空间省去，在curRow当做lastRow，在curRow完成计算

	在原数组计算，只需要记住上一个元素的值即可

	[1,3,3,1]

	=> [1,3,3,1,1] pre = 1, j = 1
	=> tmp = cur[j]: 3, cur[j] = cur[j] + pre = 4, pre = tmp: 3
	=> [1,4,3,1,1] pre = 3, j = 2
	=> [1, 4, 6, 4, 1]
*/
func getRow_v1(rowIndex int) []int {
	cur := make([]int, rowIndex + 1, rowIndex + 1)
	for i := 0; i <= rowIndex; i++ {
		cur[i] = 1
		pre := 1
		for j := 1; j < i; j++ {
			tmp := cur[j]
			cur[j] = cur[j] + pre
			pre = tmp
		}
	}
	return cur
}