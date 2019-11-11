package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1,4,3,2}))
}

/*
	本题比较简单，思路是对数组先进行排序，将奇数的数相加就是最大的值

	为了理解这种方法，让我们从不同的角度来看待问题。我们需要形成数组元​​素的配对，使得这种配对中最小的总和最大。因此，我们可以查看选择配对中最小值的操作，比如 (a,b)(a,b) 可能会产生的最大损失 a-ba−b (如果 a > ba>b)。

	如果这类配对产生的总损失最小化，那么总金额现在将达到最大值。只有当为配对选择的数字比数组的其他元素更接近彼此时，才有可能将每个配对中的损失最小化。

	考虑到这一点，我们可以对给定数组的元素进行排序，并直接按排序顺序形成元素的配对。这将导致元素的配对，它们之间的差异最小，从而导致所需总和的最大化。

	作者：LeetCode
	链接：https://leetcode-cn.com/problems/array-partition-i/solution/shu-zu-chai-fen-i-by-leetcode/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)	
	sum := 0
	for i := 0; i < len(nums) - 1; i += 2 {
		sum += nums[i]
	}
	return sum
}