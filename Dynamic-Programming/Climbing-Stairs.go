package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs1(5))
}

/*
	动态规划

	将当前问题转化为子问题的求解

	关于本题的直接讨论见文章： https://juejin.im/post/5a29d52cf265da43333e4da7
*/

// 使用递归来解决，很简单
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
  return climbStairs(n - 1) + climbStairs(n - 2)
}

// 使用迭代
func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	list := make([]int, n, n)
	list[0] = 1
	list[1] = 2
	for i := 2; i < n; i++ {
		list[i] = list[i - 1] + list[i - 2]
	}
	return list[n - 1]
}