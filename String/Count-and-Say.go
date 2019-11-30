package main

import (
	"fmt"
	// "strconv"
)

func main() {
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(6))

	fmt.Println(countAndSay_v1(3))
	fmt.Println(countAndSay_v1(6))
}

/*
	leetcode的题目描述看不懂，导致题目无从下手，可以参考下面的链接的题目描述

	https://zhuanlan.zhihu.com/p/34300515

	这个题目有点像斐波那契数列计算，可以使用的递归来做

	1  1
	2  11
	3  21
	4  1211
	5  111221
*/
// 递归的实现
func countAndSay(n int) string {
  if n == 1 {
    return "1"
	}
  preStr := countAndSay(n - 1)
	curByte, count, curStr := preStr[0], 1, ""
  for i := 1; i < len(preStr); i++ {
    if curByte == preStr[i] {
      count++
      continue
    }
    curStr += fmt.Sprintf("%d%s", count, string(curByte))
		curByte = preStr[i]
		count = 1
	}
	curStr += fmt.Sprintf("%d%s", count, string(curByte))
  return curStr
}

// 递推的实现，注意不断申请局部变量会影响执行速度（从leetcode的执行耗时观察）
func countAndSay_v1(n int) string {
	preStr := "1"
	curByte, count, curStr := preStr[0], 1, ""
	for i := 2; i <= n; i++ {
		curByte, count, curStr = preStr[0], 1, ""
		for j := 1; j < len(preStr); j++ {
			if curByte == preStr[j] {
				count++
				continue
			}
			curStr += fmt.Sprintf("%d%s", count, string(curByte))
			curByte = preStr[j]
			count = 1
		}
		curStr += fmt.Sprintf("%d%s", count, string(curByte))
		preStr = curStr
	}
	return preStr
}
