package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}

// API
func strStr_1(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}

/*
	暴力破解
	并假设现在文本串S匹配到 i 位置，模式串P匹配到 j 位置，则有：
	* 如果当前字符匹配成功（即S[i] == P[j]），则i++，j++，继续匹配下一个字符；
	* 如果失配（即S[i]! = P[j]），令i = i - (j - 1)，j = 0。相当于每次匹配失败时，i 回溯，j 被置为0。
*/
func strStr_2(haystack string, needle string) int {
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j < len(needle) {
		return -1
	}
	return i - j
}

/*
	KMP算法
	基于暴力破解中的i的回退的改进，在i进行回退的时候，保证让i回退到有效的位置，减少匹配的次数
	设法利用这个已知信息，不要把"搜索位置"移回已经比较过的位置，继续把它向后移，这样就提高了效率。

	KMP算法通过计算部分匹配值来实现有效回退
	"部分匹配值"就是"前缀"和"后缀"的最长的共有元素的长度

	例如：
	"ABCDA"的前缀为[A, AB, ABC, ABCD]，后缀为[BCDA, CDA, DA, A]，共有元素为"A"，长度为1;

	首先有一个存储，来存储计算的子字符串的部分匹配值的长度

	还有一个函数，来计算部分匹配值

	之后改造一下暴力破解的函数即可

	TBD：
*/
func strStr(haystack string, needle string) int {
	cache := make(map[int]int, 0) // 字符串的长度: 对应的部分匹配值+1，这里和默认值0区分开来，0需要走一次计算 1表示已经计算过了，但是匹配值是0

	computePartialMatch := func (s string) int {
		if len(s) == 1 {
			return 0
		}
		if cache[len(s)] > 0 {
			return cache[len(s)] - 1
		}
	
	}

	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i + 1 - computePartialMatch(string(needle[:j]))
			j = 0
		}
	}

	if j < len(needle) {
		return -1
	}

	return i - j
}





