package commonsort

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

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l int, r int) int {
	// 选择基准值
	m := nums[l]
	s := l
	l++

	for l <= r {
		// 注意这里的下表判断应该优先，不然会有越界的问题
		for l <= r && nums[l] <= m {
			l++
		}
		for l <= r && nums[r] > m {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}

	nums[r], nums[s] = nums[s], nums[r]

	return r
}
