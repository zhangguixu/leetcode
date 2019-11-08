package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber_v2([]int{2, 2, 1}))
	fmt.Println(singleNumber_v2([]int{4, 1, 2, 1, 2}))
}

/*
	借助hashtable的思路：

	时间复杂度： O(n)
	空间复杂度： O(n)

	代码中的target的作用就是避免再遍历一次hashtable，举例说明

	[1,1,3,2,2]

	target: 1 - 1 + 3 + 2 -2 = 3
*/
func singleNumber(nums []int) int {
	countMap := make(map[int]int, len(nums))
	target := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if countMap[n] == 0 {
			countMap[n] = 1
			target += n
		} else {
			target -= n
		}
	}
	return target
}

/*
	trick: 异或运算

	位运算处理： https://lihaoquan.me/2018/1/1/bit-operator.html
	https://zh.wikipedia.org/wiki/%E4%BD%8D%E6%93%8D%E4%BD%9C

	思路其实跟hashtable的一样的，能够将相同的数字抵消，（因为相同数字的出现的次数是2，为偶数）

	利用异或运算的特点

	a ^ a  = 0
	0 ^ b = b

	a ^ a ^ b = b
*/
func singleNumber_v1(nums []int) int {
	r := 0
	for _, n := range nums {
		r ^= n
	}
	return r
}

/*
	排序之后，再遍历

	时间复杂度： O(n) ~ O(nlogn)
	空间复杂度：O(i)
*/
type List []int

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	return l[i] < l[j]
}

func singleNumber_v2(nums []int) int {
	sort.Sort(List(nums))
	i := 0
	for i < len(nums)-1 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		i += 2
	}
	return nums[i]
}
