package main

import (
	"fmt"
	"math/rand"
	// "time"
)


/*
	这道题考察就是洗牌（shuffle）算法

	参考： https://blog.csdn.net/qq_26399665/article/details/79831490

	此题目是已知数组长度的，因此可以使用

	- Fisher-Yates Shuffle算法
	- Knuth-Durstenfeld Shuffle 算法 （改进版的Fisher-Yates Shuffle算法）

	
	
	l有一点需要注意的是，leetcode是如何评判题目的正确性：

	go中的rand，随机数从资源生成。包水平的函数都使用的默认的公共资源。该资源会在程序每次运行时都产生确定的序列。
	如果需要每次运行产生不同的序列，应使用Seed函数进行初始化。默认资源可以安全的用于多go程并发。

	因此这道题目中，
	
	- 不能自己指定rand的Seed，否则会一直是wrong answer
	- 随机选择的数必须要从0开始交换起，否者也无法跟答案一样

*/

type Solution struct {
	OriginNums []int
}


func Constructor(nums []int) Solution {
    return Solution{
		OriginNums: nums,
	}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.OriginNums
}

func (this *Solution) Random(min, max int) int {
	return min + int(rand.Float64() * float64(max - min))
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, 0, len(this.OriginNums))
	nums = append(nums, this.OriginNums...)

	n := len(nums)

	for i := 0; i < n; i++ {
		randIndex := this.Random(i, n)
		nums[i], nums[randIndex] = nums[randIndex], nums[i]
	}

	return nums
}

func main() {
	solution := Constructor([]int{1,2,3,4,5,6})
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
}
