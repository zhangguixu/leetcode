package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

/*
	本题

	求面积的公式是

	area = (j - i) * min(a[j], a[i]) j > i

	若要面积最大，思路是先从 j - i 的最大值开始，即使用双指针指向head和tail，保证j - i的最大

	那这里指针该如何移动呢?

	我们知道移动指针之后，tail - head的值是固定的，max(a[j], a[i])决定了计算所得面积的上限就越大，因此保留两者之间的最大值。

	移动指针就必须从短值进行移动，也就是若

		a[i] < a[j], i++,
		a[j] > a[i], j--
	
	总结：

	1）这道题虽然还是双指针法，但是关键在于对于求面积公式的理解，本题未能完成，是因为思路在于缩小min(a[j], a[i])的距离，而不是在于j - i
	2）指针移动方向的理解
*/
func maxArea(height []int) int {
	head, tail, max := 0, len(height) - 1, 0
	for head < tail {
		min := height[head]
		width := tail - head
		if height[tail] < min {
			min = height[tail]
			tail--
		} else {
			head++
		}
		area := min * width
		if max < area {
			max = area
		}
	}
	return max
}

