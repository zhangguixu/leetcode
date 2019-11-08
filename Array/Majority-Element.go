package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

/*
	hashtable

	时间复杂度为O(n)
	空间复杂度为O(n)

	注意使用除法判断会有精度的问题，导致判断不准，如长度是7 / 2 = 3，会导致示例case无法通过
*/
func majorityElement(nums []int) int {
	countMap := make(map[int]int, 0)
	total := len(nums)

	for i := 0; i < total; i++ {
		countMap[nums[i]]++
		if countMap[nums[i]] * 2 >= total {
			return nums[i]
		}
	}

	return 0
}

// nlog(n)的做法，先排序，由于是众数的定义，因此下标为[n / 2]或者[n / 2 + 1]就是众数
/*
 这里使用的go的sort进行排序，从leetcode的运行结果来看，效率并没有提升。
 sort包中实现了３种基本的排序算法：插入排序．快排和堆排序．
 和其他语言中一样，这三种方式都是不公开的，他们只在sort包内部使用．
 所以用户在使用sort包进行排序时无需考虑使用那种排序方式，
 sort.Interface定义的三个方法：获取数据集合长度的Len()方法、
 比较两个元素大小的Less()方法和交换两个元素位置的Swap()方法，就可以顺利对数据集合进行排序。
 sort包会根据实际数据自动选择高效的排序算法。
*/
func majorityElement_v1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
