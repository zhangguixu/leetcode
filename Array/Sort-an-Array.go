package main

import "fmt"

func main() {
	nums := []int{2,1,3,5,5}
	quickSort(nums)
	fmt.Println(nums)
}

/*
	排序算法练习
*/


/*
	快排的思路

	选择一个基准
	比较，将小于基准的放在左边，大于基准数的放在右边
	然后再排序的左边数组，右边的数组

	1）从数列中挑出一个元素，称为"基准"（pivot）。
	2）重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（相同的数可以到任何一边）。在这个分区结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。
	3）递归地（recursively）把小于基准值元素的子数列和大于基准值元素的子数列排序。

	快排是不稳定的排序，在极端情况下，时间复杂度为O(n * n)
	一般情况下是O(nlogN)

	实现默认选择第一个数为基准值

	快排有一些可以优化的点：

	1）当数列近乎有序的时，由于每次选取的都是第一个数，所以造成数列分割的极其不等，此时快排蜕化成 O(n^2) 的算法， 此时只要随机选取基准点即可
	2）当数列中包含大量的重复元素的时候，这一版的代码也会造成"分割不等“的问题，此时需要将重复元素均匀的分散的自数列旁
	3）使用三路快排
*/
func quickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	quickSortHelper(nums, 0, len(nums) - 1)
}

func quickSortHelper(nums []int, l, h int) {
	if (l >= h) {
		return 
	}
	p := partition(nums, l, h)
	quickSortHelper(nums, l, p - 1)
	quickSortHelper(nums, p + 1, h)
}

func partition(nums []int, l, h int) int {
	v, i, j := nums[l], l + 1, h
	for i <= j {
		for i <= j && nums[i] <= v {
			i++
		}
		for i <= j && nums[j] >= v {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

