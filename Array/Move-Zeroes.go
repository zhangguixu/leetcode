package main

import "fmt"

func main() {
	moveZeroes_v1([]int{0, 1, 0, 3, 12})
	moveZeroes_v1([]int{0, 1, 0, 3, 12, 0, 0, 0})
	moveZeroes_v1([]int{1, 1, 1, 3, 12, 0, 0, 0})
}

/*
	如果没有in-place的要求的话，可以借助另外一个数组，这样也很容易实现

	时间复杂度是O(n)
	空间复杂度是O(n)
*/

// 实现略，比较简单

/*
	in-place的技巧：数组元素的交换

	若当前的值为0，在后面找到一个非0的元素进行交换

	这种实现，时间复杂度： 
		n + n - 1 + n - 2 + .. +1 = n(n-1)/ 2,即O(n2)
	空间复杂度：O(1)
*/
func moveZeroes(nums []int) {

	l := len(nums)

	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			continue
		}
		allZero := true
		for j := i + 1; j < l; j++ {
			if nums[j] != 0 {
				allZero = false
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
		if allZero {
			break
		}
	}

	fmt.Println(nums)
}

/*
	根据solution的第三种方法，数组元素交换可以使用two-pointer来实现，这样的代码更加的简洁，实现的思路

	数组                  cur         nonZeroIndex

	[0, 1, 0, 3, 12]      0              0

	[0, 1, 0, 3, 12]      1              0                 交换

	[1, 0, 0, 3, 12]      2              1 

	[1, 0, 0, 3, 12]      3              1                 交换

	[1, 3, 0, 0, 12]      4              2                 交换

	[1, 3, 12, 0, 0]      5              3                 数组循环结束

	时间复杂度为： O(n)
	空间复杂度：O(1)
*/
func moveZeroes_v1(nums []int) {
	nonZeroIndex := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[nonZeroIndex], nums[cur] = nums[cur], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}
	fmt.Println(nums)
}

/*
	in-place的技巧：数组元素的覆盖
	
	two-pointer

	时间复杂度：O(n)

	空间复杂度：O(1)

*/
func moveZeroes_v2(nums []int) {
	lastFoundZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastFoundZero] = nums[i]
			lastFoundZero++
		}
	}
	for lastFoundZero < len(nums) {
		nums[lastFoundZero] = 0
		lastFoundZero++
	}
	fmt.Println(nums)
}
