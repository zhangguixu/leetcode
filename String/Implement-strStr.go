package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(strStr("hello", "ll"))
	// fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(computeNext("ABCDABD"))
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
func strStr(haystack string, needle string) int {
  var i, k, j int
  for i = 0; i <= len(haystack) - len(needle); i++ {
    k = i
    for j = 0; j < len(needle); j++ {
      if haystack[k] != needle[j] {
        break
      } 
      k++
    }
    if k - i == len(needle) {
      return i
    }
  }
  return -1
}

/*
	KMP算法：目前讲的最好就是https://blog.csdn.net/v_july_v/article/details/7041827，注意一些细节

	需要注意的细节

	- 移动的下标是j
	- next数组表示是失配的时候，j移动下一个位置
	- next数组是PMT表整体右移一位，然后初值赋为-1
	- PMT（Partial Match Table) 前缀后缀最长公共元素长度
	- 求next数组的计算过程中，需要用到模式串的自我匹配，这个是理解next数组，理解整个KMP的关键，注意这里的下标

	到这里就很好理解，求next[j]，其实就是计算P[0...,j -1]的前缀后缀最长的公共元素长度

	这里采用的是递推的方式来计算next数据

	next[j] = k: 意味着：p[0, k - 1] = p[j - k, j - 1] 

	已知next[j] = k，求next[j + 1]

	若p[k] == p[j]，那么，next[j + 1] = next[j] + 1 
	若p[k] != p[j]，那就等于
	p[0, k] 与 p[j - k, j]在p[k]的时候失配了，因为next[k]已经是计算出来了，在失配的时候（请将其视为两个字符串在匹配，p[j -k, j] 是字符串，而p[0,k]是模式串）
	这里很玄妙，也是我觉得KMP算法最难理解的地方，跟一开始所要进行字符的模式匹配是类似的问题，而且这个时候next[k]是已知的，而且我们仅需要找到一个p[next[k]] == p[j]即可，不需要模式串的完全匹配

	后面会在总结一篇文章
*/
//KMP算法：计算next数组
func computeNext(str string) []int {
	next := make([]int, len(str), len(str))
	next[0] = -1
	var k int
	for i := 1; i < len(str); i++ {
		k = next[i - 1]
		for {
			if k == -1 || str[i - 1] == str[k] {
				next[i] = k + 1
				break
			}
			k = next[k]
		}
	}
	return next
}