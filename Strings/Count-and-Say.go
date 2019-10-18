package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(5))
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
func countAndSay(n int) string {
	return say(n)
}

func say(n int) string {
	if n == 1 {
		return "1"
	}
	var curStr string
	preStr := say(n - 1)
	sayStr := preStr[0]
	count := 1
	for i := 1; i < len(preStr); i++ {
		if sayStr == preStr[i] {
			count++
			continue
		}
		curStr += strconv.Itoa(count) + string(sayStr)
		sayStr = preStr[i]
		count = 1
	}
	curStr += strconv.Itoa(count) + string(sayStr)
	return curStr
}