package main

import "fmt"

func main() {
	rotate_v3([]int{1, 2, 3, 4, 5, 6, 7}, 1)
	rotate_v3([]int{1, 2, 3, 4, 5, 6, 7}, 2)
	rotate_v3([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate_v3([]int{1, 2}, 3)
	rotate_v3([]int{1, 2, 3}, 4)
	rotate_v3([]int{1, 2, 3, 4, 5, 6}, 2)

	// visitor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 8)

}

/*
	最简单的实现方式，就是借助额外的空间，

	例如[1,2,3,4,5,6,7] k = 3

	申请一个新的数组，并且

	[5,6,7]

	然后append操作，将原数组中[1,2,3,4]添加到新数组后面

	最后再拷贝，将新数组的元素覆盖原数组的

	这里还需要注意的是：

	k是有可能大于n的，当k == n的时候，数组刚刚旋转一周，元素的位置都不变，
	因此在代码实现的时候，都会先对k进行处理：

		k = k % n

	算法复杂度分析

  Time complextiy : O(n)
	Space complexity : O(n)
*/
func rotate(nums []int, k int) {
	// k有可能大于数组的长度的
	l := len(nums) - k%len(nums)
	list := nums[l:]
	for i := 0; i < l; i++ {
		list = append(list, nums[i])
	}
	copy(nums, list)
}

/*
	根据solution提供的方法，Cyclic Replacements
	这道题的建立在一个知识点，见visitor函数
*/
func rotate_v1(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			fmt.Println(current, next, nums, start)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			if start == current {
				break
			}
		}
	}
	fmt.Println(nums)
}

/*
	以上的代码可以调整为：
*/
func rotate_v2(nums []int, k int) {
	k = k % len(nums)
	count := 0
	start := 0
	for count < len(nums) {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			fmt.Println(current, next, nums, start)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			if start == current {
				break
			}
		}
		start++
	}
	fmt.Println(nums)
}

/*
	当n与k的最大公约数大于1的时候，从起点start，以k为步长，回来起点start的时候，无法访问到所有的元素，需要继续start++，直到访问次数为n，就可以访问且仅一次访问到所有的元素

	visitor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 6)

	visitor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 8)
*/
func visitor(nums []int, start int, k int) {
	current := start
	for {
		fmt.Printf(" %d ", nums[current])
		next := (current + k) % len(nums)
		current = next
		if start == current {
			break
		}
	}
	fmt.Printf("\n")
}

/*
	solution的第三种解法，就是通过固定的三次reverse来实现数组的旋转

	[1,2,3,4,5]  k = 3

	整个数组反转

	[5,4,3,2,1]

	数组的前k个元素反转，

	[3,4,5,2,1]

	数组后半部分 n - k个元素反转

	[3,4,5,1,2]

	在disucss中看到谈论如何会想到这种解法，看到有一条comment说的特别好

	My understanding is the core concept of this algorithm is the reverse of reverse is original order.
	WIth this concept in mind, this method becomes very straightforward.

	观察数组，旋转的效果跟reverse的效果有点类型，都是把两半部分调换，这是一个思路切入点

	在discuss看到的极致的抽象，很有意思

	------>-->
	<------<--
	-->------>
*/
func rotate_v3(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[start], nums[end]
		start++
		end--
	}
}
